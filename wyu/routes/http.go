package routes

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"wyu/app/http/controllers"
	"wyu/configs"
	"wyu/modules"
)

type http struct {
	r *gin.Engine
}

func NewHttp(r *gin.Engine) *http {
	return &http{r:r}
}

func (h *http) HttpRoutes() {
	//h.r.Use(middleware.M(), middleware.MSession())

	for key, val := range h.toHttp() {
		sp := strings.Split(key, ",")
		switch strings.ToLower(sp[1]) {
		case "get":
			h.r.GET (configs.WYuRouteHttp[sp[0]], val ...)
			continue

		case "post":
			h.r.POST(configs.WYuRouteHttp[sp[0]], val ...)
			continue

		default:
			continue

		}
	}
}

func (h *http) HttpFuncMap() template.FuncMap {
	return template.FuncMap{
		"T": modules.I18nT,
	}
}

func (h *http) toHttp() map[string][]gin.HandlerFunc {
	var cIndex controllers.Index = controllers.NewIndexController()

	return map[string][]gin.HandlerFunc{
		"RHIndexTest,get": []gin.HandlerFunc{cIndex.Test},
	}
}