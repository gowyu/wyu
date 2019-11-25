package controllers

import (
	services2 "wyu/app/server/apis/services"
)

type controller struct {
	srv *services2.Service
}

func NewController() *controller {
	return &controller{}
}
