package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
	"wyu/app/exceptions"
	cApis "wyu/app/server/apis/controllers"
	cHttp "wyu/app/server/http/controllers"
	"wyu/configs"
	"wyu/modules"
)

const SP string = "->"

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
			"HIndex" +SP+ "get" +SP+ "/": []gin.HandlerFunc{HttpToIndex.Index},
			"HIndexTests" +SP+ "get" +SP+ "/tests": []gin.HandlerFunc{HttpToIndex.Tests},
			"HIndexHtmls" +SP+ "get" +SP+ "/htmls": []gin.HandlerFunc{HttpToIndex.Htmls},
			"HIndexCache" +SP+ "get" +SP+ "/cache": []gin.HandlerFunc{HttpToIndex.Cache},
		},
		"APIS": map[string][]gin.HandlerFunc{
			"ATest" +SP+ "get" +SP+ "/test": []gin.HandlerFunc{ApisToTests.Tests},
		},
	}
}

func To(r *gin.Engine) {
	/**
	 * No Route To Redirect
	**/
	r.NoRoute(exceptions.NewExceptions().NoRoute)
	
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
	var tplFunc template.FuncMap = template.FuncMap{}

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

func do(g *gin.RouterGroup, toFunc map[string][]gin.HandlerFunc) {
	for route, ctrl := range toFunc {
		Y := strings.Split(route, SP)

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