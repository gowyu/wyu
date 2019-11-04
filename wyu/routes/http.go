package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
	"wyu/app/http/controllers"
	"wyu/app/middleware"
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
	h.r.Use(middleware.M())

	for key, val := range h.toHttp() {
		z := strings.Split(key, ",")

		x, ok := configs.WYuRouteHttp[z[0]];
		if ok == false {
			continue
		}

		switch strings.ToLower(z[1]) {
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
	return template.FuncMap {
		"T": modules.I18nT,
		"U": middleware.TviewURL,
	}
}

func (h *http) toHttp() map[string][]gin.HandlerFunc {
	var cIndex controllers.Index = controllers.NewIndexController()

	return map[string][]gin.HandlerFunc{
		"RH___g1,get":[]gin.HandlerFunc{cIndex.Index},
		"RH___g2,get":[]gin.HandlerFunc{cIndex.Tests},
		"RH___g3,get":[]gin.HandlerFunc{cIndex.Htmls},
	}
}