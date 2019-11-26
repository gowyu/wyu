package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"wyu/app/middleware"
	"wyu/modules"
)

type http struct {

}

func (h *http) Tag() string {
	return "HTTP"
}

func (h *http) Put(r *gin.Engine, toFunc map[string][]gin.HandlerFunc) {
	do(r.Group("", middleware.M()), toFunc)
}

func (h *http) ToFunc() template.FuncMap {
	return template.FuncMap{
		"T": modules.I18nT,
		"U": middleware.TviewURL,
	}
}