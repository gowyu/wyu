package subscribe

import (
	"wyu/app/console/provider"
	"wyu/modules"
)

type rd struct {}

func (subscribed *rd) Subscribe(channels ...string) {
	subscribe := modules.InstanceRedis().Engine().Subscribe(channels ...)

	_, err := subscribe.Receive()
	if err != nil {
		return
	}

	srv := provider.NewSubscribe()
	for msg := range subscribe.Channel() {
		srv.Do(msg)
	}
}


