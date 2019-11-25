package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	ctr *controller
}

func NewApi() *Api {
	return &Api{
		ctr: NewController(),
	}
}

func (c *Api) Tests(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg":"Json Testing ..."})
}
