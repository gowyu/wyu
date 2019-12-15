package services

import (
	wYu "wyu/app/models/wYu_model"
	"wyu/configs"
)

type IndexService interface {
	Test(cols []string, query interface{}, args ...interface{}) []wYu.Tests
}

type indexService struct {
	srv *Service
	mTest *wYu.TestsModel
}

var (
	_ Services     = &indexService{}
	_ IndexService = &indexService{}
)

func NewIndexService() *indexService {
	return &indexService{
		srv:   NewService(),
		mTest: wYu.NewTestsModel(),
	}
}

func (s *indexService) Paginator(strPage string, strSize string) map[string]interface{} {
	return nil
}

func (s *indexService) Test(cols []string, query interface{}, args ...interface{}) []wYu.Tests {
	var dbInitialized configs.MdbInitialized

	if cols != nil {
		dbInitialized.Columns = cols
	}

	if query != nil && len(args) > 0 {
		dbInitialized.Query = query
		dbInitialized.QueryArgs = args
	}

	wYuTests, _ := s.mTest.FetchAllByCondition(dbInitialized)
	return wYuTests
}
