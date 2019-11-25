package controllers

import (
	"github.com/gin-gonic/gin"
	services2 "wyu/app/server/http/services"
)

type Index struct {
	ctr      *controller
	srv      services2.Services
	srvIndex services2.IndexService
}

func NewIndexController() *Index {
	return &Index{
		ctr:      NewController(),
		srv:      services2.NewIndexService(),
		srvIndex: services2.NewIndexService(),
	}
}

func (c *Index) Index(gc *gin.Context) {
	c.ctr.To(gc, gin.H{"msg": "test success index"})
}

func (c *Index) Tests(gc *gin.Context) {
	c.ctr.srv.Cache.Publish("test","test publish success ...")

	c.ctr.To(
		gc,
		gin.H{
			"msg": "test success testing ...",
			"txt": c.srvIndex.Test(nil, nil, nil),
			"rds": c.ctr.srv.Cache.Get("test").Val(),
		},
	)
}

func (c *Index) Htmls(gc *gin.Context) {
	c.ctr.srv.Cache.Publish("test","test publish HTML success ...")
	c.ctr.To(gc, gin.H{"msg": "test success html ..."}, "index.html")
}

func (c *Index) Cache(gc *gin.Context) {
	c.ctr.To(
		gc,
		gin.H{
			"rds": c.ctr.srv.Cache.Get("test").Val(),
			"rdc": c.ctr.srv.Cache.Get("test").Args(),
		},
	)
}
