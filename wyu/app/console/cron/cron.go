package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"wyu/modules"
)

func init() {
	if modules.Env == nil {
		panic("get env config error in console/cron/cron.go")
	}

	if modules.Env.GET("YuCronTab", false).(bool) {
		new(cronTab).do()
	}
}

type cronTab struct {

}

func (cTab *cronTab) do() {
	i := 0
	c := cron.New()

	c.AddFunc("*/5 * * * * ?", func () {
		i++
		fmt.Println("cron running ...", i)
	})

	c.Start()
	defer c.Stop()

	select {

	}
}