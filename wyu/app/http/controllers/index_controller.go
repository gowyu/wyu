package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index interface {
	Test(c *gin.Context)
}

type index struct {
	ctrl Controller
}

var _ Index = &index{}

func NewIndexController() *index {
	return &index{
		ctrl: NewController(),
	}
}

func (c *index) Test(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"msg":"test success"})
}
