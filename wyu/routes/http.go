package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
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
		H := strings.Split(key, ",")

		x, ok := configs.WYuRouteHttp[H[0]];
		if ok == false {
			continue
		}

		switch strings.ToLower(H[1]) {
		case "get":
			h.r.GET (x, val ...)
			continue

		case "post":
			h.r.POST(x, val ...)
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
		"RHIndex,get": []gin.HandlerFunc{cIndex.Index},
		"RHIndexTest,get": []gin.HandlerFunc{cIndex.Test},
	}
}