package controllers

import (
	"github.com/gin-gonic/gin"
	"wyu/app/http/services"
)

type Index interface {
	Index(c *gin.Context)
	Tests(c *gin.Context)
	Htmls(c *gin.Context)
}

type index struct {
	ctrl *controller
	srv services.Services
	srvIndex services.IndexService
}

var _ Index = &index{}

func NewIndexController() *index {
	return &index{
		ctrl: NewController(),
		srv: services.NewIndexService(),
		srvIndex: services.NewIndexService(),
	}
}

func (c *index) Index(gc *gin.Context) {
	c.ctrl.To(gc, gin.H{"msg":"test success index"})
}

func (c *index) Tests(gc *gin.Context) {
	c.ctrl.To(gc, gin.H{"msg":"test success testing ..."})
}

func (c *index) Htmls(gc *gin.Context) {
	c.ctrl.To(gc, gin.H{"msg":"test success html ..."}, "index.html")
}
