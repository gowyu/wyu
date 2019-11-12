package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

type apis struct {

}

func (h *apis) Tag() string {
	return "APIS"
}

func (h *apis) Put(r *gin.Engine, toFunc map[string][]gin.HandlerFunc) {
	ToHandle(r, toFunc)
}

func (h *apis) ToFunc() template.FuncMap {
	return nil
}
