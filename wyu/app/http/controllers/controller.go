package controllers

type Controller interface {

}

type controller struct {

}

var _ Controller = &controller{}

func NewController() *controller {
	return &controller{}
}
