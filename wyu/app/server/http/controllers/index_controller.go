package controllers

import (
	"github.com/gin-gonic/gin"
	"wyu/app/server/http/services"
)

type Index struct {
	//wxc *modules.WeChat
	ctr *controller
	srvIndex *services.IndexSrv
}

func NewIndexController() *Index {
	return &Index{
		//wxc: modules.NewWeChat(), // wechat
		ctr: NewController(),
		srvIndex: services.NewIndexService(),
	}
}

func (c *Index) Index(gc *gin.Context) {
	c.ctr.To(gc, gin.H{"txt": "Test WeiXin Login"}, "index.html")
}

func (c *Index) Tests(gc *gin.Context) {
	//fmt.Println(c.srvIndex.SendMail().Error())
	c.srvIndex.Token()
	c.ctr.To(
		gc, gin.H{
			//"msg": "test success testing ...",
			//"txt": c.srvIndex.Test([]string{"name"}, nil, nil),
			//"tid": c.srvIndex.TestById(nil, "id=?", "2"),
			//"num": c.srvIndex.Nums(),
			"exe": c.ctr.srv.Parent.R.Exists("test").Val(),
			"txt": c.srvIndex.GetCache("test"),
			//"join": c.srvIndex.TestToTest(nil, nil, nil),
			"tests": c.srvIndex.Tables(3, 0),
			"tests_paginator": c.srvIndex.Paginator("1", "3"),
		},
	)
}

func (c *Index) Htmls(gc *gin.Context) {
	c.ctr.To(gc, gin.H{"msg": "test success html ..."}, "index.html")
}

func (c *Index) Cache(gc *gin.Context) {
	/**
	 * Todo: Subscribe & Publish
	 */
	c.ctr.srv.Parent.R.Publish("service","test publish success1 ...")

	c.ctr.To(
		gc, gin.H{
			"rds": "test publish ...",
		},
	)
}
