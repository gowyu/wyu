package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	ctrl *controller
}

func NewApi() *Api {
	return &Api{
		ctrl: NewController(),
	}
}

func (c *Api) Tests(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg":"Json Testing ..."})
}
