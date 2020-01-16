package services

import (
	"fmt"
	"wyu/app/models"
	wYu "wyu/app/models/wYu_model"
	"wyu/configs"
)

type IndexSrv struct {
	srv *Service
	mTest *wYu.TestsModel
}

func NewIndexService() *IndexSrv {
	return &IndexSrv {
		srv: NewService(),
		mTest: wYu.NewTestsModel(),
	}
}

func (s *IndexSrv) GetCache(key string) (str string) {
	return s.srv.Parent.R.Get(key).Val()
}

func (s *IndexSrv) Tables(limit int, start ...int) (wYuTests []wYu.Tests) {
	var dbInitialized configs.MdbInitialized

	dbInitialized.Limit = limit
	dbInitialized.Start = start

	wYuTests, _ = s.mTest.Tables(dbInitialized)
	return
}

func (s *IndexSrv) Paginator(strPage string, strSize string) (paginator map[string]interface{}) {
	page, size := s.srv.PaginatorParams(strPage, strSize)
	nums, err := s.mTest.Nums()
	if err != nil {
		nums = 0
	}

	paginator = map[string]interface{}{
		"tests": s.Tables(size, (page-1)*size),
		"paginator": s.srv.Paginator(page, int(nums), size),
	}
	return
}

func (s *IndexSrv) AddTransaction() (ok bool) {
	err := s.mTest.AddTransaction()
	if err == nil {
		ok = true
	}

	return
}

func (s *IndexSrv) Add() (ok bool) {
	_, err := s.mTest.Add(&wYu.Tests{
		Name: "Add_Test_Two",
	})

	if err == nil {
		ok = true
	}

	return
}

func (s *IndexSrv) Upd(query interface{}, args ...interface{}) (ok bool) {
	var dbInitialized configs.MdbInitialized

	if query == nil || len(args) < 1 {
		return
	}

	dbInitialized.Query = query
	dbInitialized.QueryArgs = args

	_, err := s.mTest.Upd(dbInitialized, &wYu.Tests{Name:"update success"})
	if err == nil {
		ok = true
	}

	return
}

func (s *IndexSrv) Del(query interface{}, args ...interface{}) (ok bool) {
	var dbInitialized configs.MdbInitialized

	if query == nil || len(args) < 1 {
		return
	}

	dbInitialized.Query = query
	dbInitialized.QueryArgs = args

	_, err := s.mTest.Del(dbInitialized)
	if err == nil {
		ok = false
	}

	return
}

func (s *IndexSrv) TestById(cols []string, query interface{}, args ...interface{}) (wYuTest *wYu.Tests) {
	var dbInitialized configs.MdbInitialized

	if cols != nil {
		dbInitialized.Columns = cols
	}

	if query == nil || len(args) < 1 {
		return
	}

	dbInitialized.Query = query
	dbInitialized.QueryArgs = args

	wYuTest, _ = s.mTest.FetchOne(dbInitialized)
	return
}

func (s *IndexSrv) Test(cols []string, query interface{}, args ...interface{}) (wYuTests []wYu.Tests) {
	var dbInitialized configs.MdbInitialized

	if cols != nil {
		dbInitialized.Columns = cols
	}

	if query != nil && len(args) > 0 {
		dbInitialized.Query = query
		dbInitialized.QueryArgs = args
	}

	wYuTests, _ = s.mTest.FetchAll(dbInitialized)
	return
}

func (s *IndexSrv) TestToTest(cols []string, query interface{}, args ...interface{}) (wYuTestToTest []wYu.TestToTest) {
	var dbInitialized configs.MdbInitialized

	if cols != nil {
		dbInitialized.Columns = cols
	}

	if query != nil && len(args) > 0 {
		dbInitialized.Query = query
		dbInitialized.QueryArgs = args
	}

	dbInitialized.Table = "tests"
	dbInitialized.Field = "tests.*, tt.*"

	dbInitialized.Joins = [][]interface{}{
		0:{"INNER", []string{"test_test","tt"}, "tests.id = tt.test_id"},
	}

	dbInitialized.OrderType = models.SelectToESC
	dbInitialized.OrderArgs = []string{"tests.id"}

	wYuTestToTest, _ = s.mTest.FetchAllJoin(dbInitialized)
	fmt.Println(wYuTestToTest)
	return
}

func (s *IndexSrv) Nums() (nums int) {
	n, err := s.mTest.Nums()
	if err != nil {
		return
	}

	nums = int(n)
	return
}

func (s IndexSrv) Subscribed(channel string, content interface{}) {
	fmt.Println("test index provider:", channel, "->", content)
}
