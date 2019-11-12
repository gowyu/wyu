package controllers

import (
	"wyu/app/apis/services"
)

type controller struct {
	srv *services.Service
}

func NewController() *controller {
	return &controller{}
}


