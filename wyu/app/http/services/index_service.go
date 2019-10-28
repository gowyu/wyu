package services

var (
	_ Services = &indexService{}
	_ IndexService = &indexService{}
)

type IndexService interface {
	Test() bool
}

type indexService struct {
	srv *Service
}

func NewIndexService() *indexService {
	return &indexService{
		srv: NewService(),
	}
}

func (s *indexService) Paginator(strPage string, strSize string) map[string]interface{} {
	return nil
}

func (s *indexService) Test() bool {
	return false
}

type testService struct {

}
