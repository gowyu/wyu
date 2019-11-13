package wYu_model

import (
	"wyu/app/models"
	"wyu/app/repositories/wYu"
	"wyu/configs"
	"wyu/modules"
)

type Tests wYu.Tests // Table Name

type TestsModel struct {
	models *models.Models
}

func NewTestsModel() *TestsModel {
	return &TestsModel{
		models: models.New(modules.InstanceClusterDB(db).Engine()),
	}
}

func (m *TestsModel) FetchAllByCondition(dbInitialized configs.MdbInitialized) []Tests {
	src := make([]Tests, 0)
	err := m.models.FetchAllByCondition(dbInitialized, &src)

	if err != nil {
		return nil
	} else {
		return src
	}
}













