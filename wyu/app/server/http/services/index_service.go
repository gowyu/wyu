package services

import (
	"fmt"
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

func (s *IndexSrv) Paginator(strPage string, strSize string) (paginator map[string]interface{}) {
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
