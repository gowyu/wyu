package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
	"wyu/app/exceptions"
	cApis "wyu/app/server/apis/controllers"
	cHttp "wyu/app/server/http/controllers"
	"wyu/configs"
	"wyu/modules"
)

type Routes interface {
	Tag() string
	Put(r *gin.Engine, toFunc map[string][]gin.HandlerFunc)
	ToFunc() template.FuncMap
}

var (
	_ Routes = new(http)
	_ Routes = new(apis)

	YuRoutes map[string]map[string][]gin.HandlerFunc
	Yu []Routes = []Routes{
		new(http),
		new(apis),
	}
)

func init() {
	if len(Yu) == 0 {
		panic("Fatal Routes")
	}

	HttpToIndex := cHttp.NewIndexController()
	ApisToTests := cApis.NewApi()

	YuRoutes = map[string]map[string][]gin.HandlerFunc{
		"HTTP": map[string][]gin.HandlerFunc{
			"HIndex" +configs.YuSep+ "get" +configs.YuSep+ "/": []gin.HandlerFunc{
				HttpToIndex.Index,
			},
			"HIndexTests" +configs.YuSep+ "get" +configs.YuSep+ "/tests": []gin.HandlerFunc{
				HttpToIndex.Tests,
			},
			"HIndexHtmls" +configs.YuSep+ "get" +configs.YuSep+ "/htmls": []gin.HandlerFunc{
				HttpToIndex.Htmls,
			},
			"HIndexCache" +configs.YuSep+ "get" +configs.YuSep+ "/cache": []gin.HandlerFunc{
				HttpToIndex.Cache,
			},
		},
		"APIS": map[string][]gin.HandlerFunc{
			"ATests" +configs.YuSep+ "get" +configs.YuSep+ "/tests": []gin.HandlerFunc{
				ApisToTests.Tests,
			},
		},
	}
}

func To(r *gin.Engine) {
	/**
	 * Todo: No Routes To Redirect
	**/
	r.NoRoute(new(exceptions.Exceptions).NoRoute)

	/**
	 * Todo: No Method To Redirect
	**/
	r.NoMethod(new(exceptions.Exceptions).NoMethod)
	
	for _, to := range Yu {
		if _, ok := YuRoutes[to.Tag()]; ok == false {
			continue
		}

		if len(YuRoutes[to.Tag()]) == 0 {
			continue
		}

		to.Put(r, YuRoutes[to.Tag()])
	}
}

func ToFunc(tpl ...interface{}) template.FuncMap {
	tplFunc := template.FuncMap{}

	for _, to := range Yu {
		if ok, _ := modules.UtilsStrContains(to.Tag(), tpl ...); ok == false {
			continue
		}

		if to.ToFunc() == nil {
			continue
		}

		for Tag, toFunc := range to.ToFunc() {
			tplFunc[Tag] = toFunc
		}
	}

	return tplFunc
}

func ToLoggerWithFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) (strLog string) {
		msg := exceptions.TxT("l^aa")

		if param.ErrorMessage != "" {
			msg = param.ErrorMessage
		}

		if param.StatusCode != 200 || param.ErrorMessage != "" {
			strLog = fmt.Sprintf(`
---------------------------------------------------------------------------------------------------
%s » %s » %s
%s » %s » %s » %s 
%s » %s
%s » %d
%s » %s
%s » %v
---------------------------------------------------------------------------------------------------
`,
				exceptions.TxT("l^ab"),
				param.ClientIP,
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				exceptions.TxT("l^ac"),
				param.Method,
				param.Request.Proto,
				param.Path,
				exceptions.TxT("l^ad"),
				param.Request.UserAgent(),
				exceptions.TxT("l^ae"),
				param.StatusCode,
				exceptions.TxT("l^af"),
				param.Latency,
				exceptions.TxT("l^ag"),
				msg,
			)
		}

		return
	})
}

func do(g *gin.RouterGroup, toFunc map[string][]gin.HandlerFunc) {
	for route, ctrl := range toFunc {
		Y := strings.Split(route, configs.YuSep)

		if len(Y) != 3 {
			continue
		}

		_, ok := configs.YuRoutes[Y[0]]
		if ok == false {
			configs.YuRoutes[Y[0]] = Y[2]
		}

		switch strings.ToLower(Y[1]) {
		case "get":
			g.GET (Y[2], ctrl ...)
			continue

		case "any":
			g.Any (Y[2], ctrl ...)
			continue

		case "post":
			g.POST(Y[2], ctrl ...)
			continue

		default:
			continue
		}
	}
}