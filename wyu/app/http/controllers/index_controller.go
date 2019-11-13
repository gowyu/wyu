package controllers

import (
	"github.com/gin-gonic/gin"
	"wyu/app/http/services"
)

type Index struct {
	ctrl *controller
	srv services.Services
	srvIndex services.IndexService
}

func NewIndexController() *Index {
	return &Index{
		ctrl: NewController(),
		srv: services.NewIndexService(),
		srvIndex: services.NewIndexService(),
	}
}

func (c *Index) Index(gc *gin.Context) {
	c.ctrl.To(gc, gin.H{"msg":"test success index"})
}

func (c *Index) Tests(gc *gin.Context) {
	c.ctrl.To(
		gc,
		gin.H{
			"msg": "test success testing ...",
			"txt": c.srvIndex.Test(nil, nil, nil),
		},
	)
}

func (c *Index) Htmls(gc *gin.Context) {
	c.ctrl.To(gc, gin.H{"msg":"test success html ..."}, "index.html")
}
