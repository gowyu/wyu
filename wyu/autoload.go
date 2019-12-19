// Copyright 2019-~ YuWenYu.  All rights reserved.
// license that can be found in the LICENSE file.

package wyu

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wyu/modules"
	"wyu/routes"

	_ "wyu/app/console/subscribe"
)

const (
	ginPort string = "8888"
	directory string = "./resources/templates/"
)

func init() {
	if modules.Env == nil {
		panic("get env configure error in autoload.go")
	}

	ad := new(autoload)
	ad.running()
}

type autoload struct {}

func (ad *autoload) running() {
	ad.ginInitialized()

	r := gin.Default()
	r = ad.ginTemplateStatic(r)

	routes.To(r)

	/**
	 * TODO: Loading Templates
	**/
	bTpl := modules.Env.GET("Temp.Status", false).(bool)
	if bTpl {
		strResources := modules.Env.GET("Temp.Resources", "").(string)
		if strResources == "" {
			log.Fatal("Templates Resources nil, Please check the configure!")
		}

		strDirViews := modules.Env.GET("Temp.DirViews", directory + "view/").(string)
		arrResources := strings.Split(strResources, ":")

		objTPL := multitemplate.NewRenderer()
		for _, skeleton := range arrResources {
			views, _ := ioutil.ReadDir(strDirViews + skeleton)
			for _, view := range views {
				arrTPL := ad.tplLoading(skeleton, view.Name())
				objTPL.AddFromFilesFuncs(view.Name(), routes.ToFunc("HTTP"), arrTPL ...)
			}
		}
		r.HTMLRender = objTPL
	}

	r.Run(":" + modules.Env.GET("App.Port", ginPort).(string))
}

func (ad *autoload) ginTemplateStatic(r *gin.Engine) *gin.Engine {
	if modules.Env.GET("Temp.StaticStatus", false).(bool) {
		static := modules.Env.GET("Temp.Static", "./resources/assets").(string)
		staticIcon := modules.Env.GET("Temp.StaticIcon", "./resources/favicon.ico").(string)

		r.Static("./assets", static)
		r.StaticFile("./favicon.ico", staticIcon)
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
		fn := dir + "/" + prefix + "_" + time.Now().Format("2006-01-02") + ".log"
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