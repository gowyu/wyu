package models

import (
	"github.com/go-xorm/xorm"
	"wyu/configs"
)

type Models struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *Models {
	return &Models {
		engine: engine,
	}
}

func (m *Models) FetchAllByCondition(dbInitialized configs.MdbInitialized, data interface{}) error {
	var xSession *xorm.Session

	if dbInitialized.Columns != nil {
		xSession = m.engine.Cols(dbInitialized.Columns ...)
	}

	if dbInitialized.Query != nil && dbInitialized.QueryArgs != nil {
		if xSession != nil {
			xSession = xSession.Where(dbInitialized.Query, dbInitialized.QueryArgs ...)
		} else {
			xSession = m.engine.Where(dbInitialized.Query, dbInitialized.QueryArgs ...)
		}
	}

	var err error

	if xSession != nil {
		err = xSession.Find(data)
	} else {
		err = m.engine.Find(data)
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}
