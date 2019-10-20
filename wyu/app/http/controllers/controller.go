package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wyu/app/http/services"
	"wyu/modules"
)

type controller struct {
	srv *services.Service
}

func NewController() *controller {
	return &controller{
		srv: services.NewService(),
	}
}

func (c *controller) Md(gc *gin.Context, arr gin.H) gin.H {
	return modules.UtilsMergeToMap(gc.MustGet("M").(gin.H), arr)
}

func (c *controller) To(gc *gin.Context, to ...interface{}) {
	var obj gin.H = gin.H{}
	if len(to) > 0 && to[0] != nil {
		obj = to[0].(gin.H)
	}

	if len(to) > 1 && to[1] != "" {
		gc.HTML(http.StatusOK, to[1].(string), gin.H{"M":c.Md(gc, obj)})
	} else {
		gc.JSON(http.StatusOK, modules.UtilsMergeToMap(obj))
	}

}
