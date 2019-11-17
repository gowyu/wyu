package modules

import (
	"fmt"
	"github.com/robfig/cron"
)

func init() {
	//NewCronTab().work()
}

type cronTab struct {

}

func NewCronTab() *cronTab {
	return &cronTab{}
}

func (cTab *cronTab) work() {
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


