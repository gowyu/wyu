package services

import (
	wYu "wyu/app/models/wYu_model"
	"wyu/configs"
)

type IndexSrv struct {
	srv *Service
	mTest *wYu.TestsModel
}

func NewIndexService() *IndexSrv {
	return &IndexSrv {
		srv:   NewService(),
		mTest: wYu.NewTestsModel(),
	}
}

func (s *IndexSrv) Paginator(strPage string, strSize string) map[string]interface{} {
	return nil
}

func (s *IndexSrv) Test(cols []string, query interface{}, args ...interface{}) []wYu.Tests {
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

func (s *IndexSrv) Nums() int {
	nums, err := s.mTest.Total()
	if err != nil {
		return 0
	}

	return int(nums)
}
