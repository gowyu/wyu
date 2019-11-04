// Copyright 2019-~ YuWenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"errors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wyu/configs"
	"wyu/modules"
	"wyu/routes"
)

const (
	ginPort string = "8888"
	directory string = "./resources/templates/"
)

func init() {
	if modules.Env == nil {
		panic(errors.New("get env configure error in autoload.go"))
	}

	wYuHttps()
	wYuInitialized()
}

func wYuHttps() {
	configs.WYuRouteHttp = map[string]string{}

	var env modules.Vipers = modules.NewVipers().Loading()
	for _, val := range env.GET("https", []interface{}{}).([]interface{}) {
		wYuSuffix := ""
		if configs.WYuSuffix != "" {
			wYuSuffix = "." + configs.WYuSuffix
		}

		H := strings.Split(val.(string), "->")

		if H[1] == "" || H[1] == "/" {
			configs.WYuRouteHttp[H[0]] = "/"
		} else {
			configs.WYuRouteHttp[H[0]] = "/" + H[1] + wYuSuffix
		}
	}
}

func wYuInitialized() {
	ad := new()
	ad.running()
}

type autoload struct {}

func new() *autoload {
	return &autoload {}
}

func (ad *autoload) running() {
	ad.ginInitialized()

	r := gin.Default()
	r = ad.ginTemplateStatic(r)

	rHttp := routes.NewHttp(r)
	rHttp.HttpRoutes()

	/**
	 * TODO: Loading Templates
	**/
	bTpl := modules.Env.GET("Temp.Status", false).(bool)
	if bTpl {
		strResources := modules.Env.GET("Temp.Resources", "").(string)
		if strResources == "" {
			panic("Templates Resources nil, Please check the configure!")
			return
		}

		strDirViews := modules.Env.GET("Temp.DirViews", directory + "view/").(string)
		arrResources := strings.Split(strResources, ":")

		objTPL := multitemplate.NewRenderer()
		for _, skeleton := range arrResources {
			views, _ := ioutil.ReadDir(strDirViews + skeleton)
			for _, view := range views {
				arrTPL := ad.tplLoading(skeleton, view.Name())
				objTPL.AddFromFilesFuncs(view.Name(), rHttp.HttpFuncMap(), arrTPL ...)
			}
		}
		r.HTMLRender = objTPL
	}

	r.Run(":" + modules.Env.GET("App.Port", ginPort).(string))
}

func (ad *autoload) ginTemplateStatic(r *gin.Engine) *gin.Engine {
	bTplStatic := modules.Env.GET("Temp.StaticStatus", false).(bool)
	if bTplStatic {
		static := modules.Env.GET("Temp.Static", "./resources/assets").(string)
		staticIcon := modules.Env.GET("Temp.StaticIcon", "./resources/favicon.ico").(string)

		r.Static("./resources/assets", static)
		r.StaticFile("./resources/favicon.ico", staticIcon)
	}

	return r
}

func (ad *autoload) ginInitialized() {
	if modules.Env.GET("Logs.Status", false).(bool) {
		dir := modules.Env.GET("Logs.Root", "./storage/logs").(string)
		_, err := os.Stat(dir)
		if err != nil {
			panic(err.Error())
			return
		}

		prefix := modules.Env.GET("Logs.Prefix", "wYu").(string)
		fn := dir + "/" + prefix + "_" + time.Now().String() + ".log"
		f, _ := os.Create(fn)
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		gin.ForceConsoleColor()
	}
}

func (ad *autoload) tplLoading(skeleton string, view string) []string {
	TplSuffix := modules.Env.GET("Temp.Suffix", "html").(string)
	dirLayout := modules.Env.GET("Temp.DirLayout", directory + "layout/").(string)

	TplLayout, err := filepath.Glob(dirLayout + skeleton + "." + TplSuffix)
	if err != nil {
		panic(err.Error())
	}

	TplShared := modules.Env.GET("Temp.DirShared", directory + "shared/").(string)
	shareds, err := filepath.Glob(TplShared + skeleton + "/" + "*.html")
	if err != nil {
		panic(err.Error())
	}

	TplViews := modules.Env.GET("Temp.DirViews", directory + "view/").(string)

	arrTPL := make([]string, 0)
	arrTPL = append(TplLayout, TplViews + skeleton + "/" + view)

	for _, shared := range shareds {
		arrTPL = append(arrTPL, shared)
	}

	return arrTPL
}