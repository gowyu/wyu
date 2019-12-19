package provider

import (
	"github.com/go-redis/redis"
	"wyu/app/server/http/services"
	"wyu/configs"
	"wyu/modules"
)

type Subscribe interface {
	Subscribed(channel string, content interface{})
}

var (
	_ Subscribe = &services.IndexSrv{}

	channels map[string]Subscribe = map[string]Subscribe{
		"service": services.NewIndexService(),
	}
)

type subscribe struct {}

func NewSubscribe() *subscribe {
	return &subscribe{}
}

func (srv *subscribe) Do(msg *redis.Message) {
	if ok, _ := modules.UtilsStrContains(msg.Channel, configs.YuSubscribe ...); ok == false {
		return
	}

	z, ok := channels[msg.Channel]
	if ok {
		z.Subscribed(msg.Channel, msg.Payload)
	}
}