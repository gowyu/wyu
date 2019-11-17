package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	NewCronTab().work()
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