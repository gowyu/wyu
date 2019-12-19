package controllers

import (
	"github.com/gin-gonic/gin"
	"wyu/app/server/http/services"
	"wyu/configs"
	"wyu/modules"
)

type Index struct {
	wxc *modules.WeChat
	ctr *controller
	srvIndex *services.IndexSrv
}

func NewIndexController() *Index {
	return &Index{
		wxc: modules.NewWeChat(),
		ctr: NewController(),
		srvIndex: services.NewIndexService(),
	}
}

func (c *Index) Index(gc *gin.Context) {
	c.ctr.To(gc, gin.H{"txt": "Test WeiXin Login"}, "index.html")
}

func (c *Index) Tests(gc *gin.Context) {
	c.ctr.To(
		gc,
		gin.H{
			"msg": "test success testing ...",
			"txt": c.srvIndex.Test(nil, nil, nil),
			"num": c.srvIndex.Nums(),
			"rds": configs.YuTest,
			"exe": c.ctr.srv.Cache.Exists("test").Val(),
		},
	)
}

func (c *Index) Htmls(gc *gin.Context) {
	c.ctr.To(gc, gin.H{"msg": "test success html ..."}, "index.html")
}

func (c *Index) Cache(gc *gin.Context) {
	c.ctr.srv.Cache.Publish("service","test publish success1 ...")
	c.ctr.To(
		gc,
		gin.H{
			"rds": "test publish ...",
		},
	)
}
