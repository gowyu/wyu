package controllers

import (
	"wyu/app/server/apis/services"
)

type controller struct {
	srv *services.Service
}

func NewController() *controller {
	return &controller{}
}
