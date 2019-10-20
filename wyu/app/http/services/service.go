package services

import "wyu/modules"

type Service struct {
	Cache modules.Rd
}

func NewService() *Service {
	return &Service{
		Cache: modules.InstanceRedis(),
	}
}
