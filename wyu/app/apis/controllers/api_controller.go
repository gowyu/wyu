package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api interface {
	Tests(c *gin.Context)
}

type api struct {
	ctrl *controller
}

var _ Api = &api{}

func NewApi() *api {
	return &api{
		ctrl: NewController(),
	}
}

func (c *api) Tests(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg":"Json Testing ..."})
}
