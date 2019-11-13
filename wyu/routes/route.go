package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
	cApis "wyu/app/apis/controllers"
	cHttp "wyu/app/http/controllers"
	"wyu/modules"
)

const SP string = "->"

type Routes interface {
	Tag() string
	Put(r *gin.Engine, toFunc map[string][]gin.HandlerFunc)
	ToFunc() template.FuncMap
}

var (
	_ Routes = &http{}
	_ Routes = &apis{}

	YuRoutes map[string]map[string][]gin.HandlerFunc
	Yu []Routes = []Routes{
		new(http),
		new(apis),
	}

)

func init() {
	if len(Yu) < 1 {
		panic("Fatal Routes")
	}

	HttpToIndex := cHttp.NewIndexController()
	ApisToTests := cApis.NewApi()

	YuRoutes = map[string]map[string][]gin.HandlerFunc{
		"HTTP": map[string][]gin.HandlerFunc{
			"/"		+SP+ "get": []gin.HandlerFunc{HttpToIndex.Index},
			"/test"	+SP+ "get": []gin.HandlerFunc{HttpToIndex.Tests},
			"/go"	+SP+ "get": []gin.HandlerFunc{HttpToIndex.Htmls},
		},
		"APIS": map[string][]gin.HandlerFunc{
			"/api/test"	+SP+ "get": []gin.HandlerFunc{ApisToTests.Tests},
		},
	}
}

func To(r *gin.Engine) {
	for _, to := range Yu {
		toTag, ok := YuRoutes[to.Tag()]
		if ok == false || len(toTag) < 1 {
			continue
		}

		to.Put(r, toTag)
	}
}

func ToFunc(r *gin.Engine, tpl ...interface{}) template.FuncMap {
	var tplFunc template.FuncMap = template.FuncMap{}
	for _, to := range Yu {
		if ok, _ := modules.UtilsStrContains(to.Tag(), tpl ...); ok == false {
			continue
		}

		toFunc := to.ToFunc()
		if toFunc == nil {
			continue
		}

		for key, val := range toFunc {
			tplFunc[key] = val
		}
	}

	return tplFunc
}

func ToHandle(r *gin.Engine, toFunc map[string][]gin.HandlerFunc) {
	for route, ctrl := range toFunc {
		Y := strings.Split(route, SP)

		switch strings.ToLower(Y[1]) {
			case "get":
				r.GET (Y[0], ctrl ...)
				continue

			case "any":
				r.Any (Y[0], ctrl ...)
				continue

			case "post":
				r.POST(Y[0], ctrl ...)
				continue

			default:
				continue
		}
	}
}