package wYu_model

import (
	"github.com/spf13/cast"
	"runtime"
	"wyu/app/exceptions"
	"wyu/app/models"
	"wyu/app/repositories/wYu"
	"wyu/configs"
	"wyu/modules"
)

type Tests wYu.Tests // Table: tests
type TestTest wYu.TestTest // Table: test_test

type TestsModel struct {
	models *models.Models
}

func NewTestsModel() *TestsModel {
	return &TestsModel{
		models: models.New(modules.InstanceClusterDB(db).Engine()),
	}
}

/**
 * Todo: Transaction
 */
func (m *TestsModel) AddTransaction() (err error) {
	T := m.models.Engine.NewSession()
	defer T.Close()

	err = T.Begin()
	if err != nil {
		return
	}

	U := &Tests{Name:"Insert Success!"}

	_, err = m.models.Insert(T, U)
	if err != nil {
		T.Rollback()
		return
	}

	if U.Id < 1 {
		_, file, line, _ := runtime.Caller(1)
		err = exceptions.Err("m^ac", file, cast.ToString(line))
		T.Rollback()
		return
	}

	_, err = m.models.Insert(T, &TestTest{TestId:U.Id, Content:"SubTest Success!"})
	if err != nil {
		T.Rollback()
		return
	}

	err = T.Commit()
	if err != nil {
		T.Rollback()
		return
	}

	return
}

/**
 * Todo: Insert
 */
func (m *TestsModel) Add(data *Tests) (i int64, err error) {
	T := m.models.Engine.NewSession()
	defer T.Close()

	i, err = m.models.Insert(T, data)
	return
}

/**
 * Todo: Update
 */
func (m *TestsModel) Upd(dbInitialized configs.MdbInitialized, data *Tests) (i int64, err error) {
	T := m.models.Engine.NewSession()
	defer T.Close()

	i, err = m.models.Update(T, dbInitialized, data)
	return
}

/**
 * Todo: Delete
 */
func (m *TestsModel) Del(dbInitialized configs.MdbInitialized) (i int64, err error) {
	T := m.models.Engine.NewSession()
	defer T.Close()

	i, err = m.models.Delete(T, dbInitialized, &Tests{})
	return
}

/**
 * Todo: SelectToOne
 */
func (m *TestsModel) FetchOne(dbInitialized configs.MdbInitialized) (src *Tests, err error) {
	dbInitialized.Types = models.SelectToOne

	src = &Tests{}
	err = m.models.Select(dbInitialized, src)

	return
}

/**
 * Todo: SelectToAll
 */
func (m *TestsModel) FetchAll(dbInitialized configs.MdbInitialized) (src []Tests, err error) {
	dbInitialized.Types = models.SelectToAll

	src = make([]Tests, 0)
	err = m.models.Select(dbInitialized, &src)

	return
}

/**
 * Todo: TotalNums
 */
func (m *TestsModel) Nums() (nums int64, err error) {
	nums, err = m.models.Total(&Tests{})
	if err != nil {
		return
	}

	return
}

