package subscribe

import (
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
		if configs.YuRoutes != nil && len(configs.YuRoutes) != 0 {
			go new(subscribe).do()
		}
	}
}

type subscribe struct {

}

func (subscribed *subscribe) do() {
	new(rd).Subscribe(configs.YuSubscribe ...)
}


