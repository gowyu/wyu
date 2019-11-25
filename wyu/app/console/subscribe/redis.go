package subscribe

import (
	"fmt"
	"wyu/modules"
)

type rd struct {

}

func (subscribed *rd) Subscribe(channels ...string) {
	subscribe := modules.InstanceRedis().Engine().Subscribe(channels ...)

	_, err := subscribe.Receive()
	if err != nil {
		return
	}

	for msg := range subscribe.Channel() {
		switch msg.Channel {
		case "test":
			fmt.Println(msg.Channel, ",", msg.Payload)
			continue
		default:
			continue
		}
	}
}


