package subscribe

import (
	"github.com/spf13/cast"
	"wyu/configs"
	"wyu/modules"
)

type Subscribe interface {
	Subscribe(channels ...string)
}

var (
	_ Subscribe = new(rd)
)

func init() {
	if modules.Env == nil {
		panic("get env config error in console/subscribe/subscribe.go")
	}

	if modules.Env.GET("YuRedisSubscribe", false).(bool) {
		go new(subscribe).do()
	}
}

type subscribe struct {

}

func (subscribed *subscribe) do() {
	new(rd).Subscribe(cast.ToStringSlice(configs.YuSubscribe) ...)
}


