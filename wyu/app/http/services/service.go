package services

import "wyu/modules"

type Service interface {

}

type service struct {
	cache modules.Rd
}

var _ Service = &service{}

func NewService() *service {
	return &service{
		cache: modules.InstanceRedis(),
	}
}
