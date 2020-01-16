package services

import (
	"wyu/app/server"
)

type Service struct {
	Parent *server.Services
}

func NewService() *Service {
	return &Service{
		Parent: server.NewServices(),
	}
}

func (srv *Service) Paginator(page int, pageNums int, pageSize ...int) (paginator map[string]interface{}) {
	paginator = srv.Parent.Paginator(page, pageNums, pageSize ...)
	return
}

func (srv *Service) PaginatorParams(strPage string, strSize string) (page int, size int) {
	page,size = srv.Parent.PaginatorParams(strPage, strSize)
	return
}
