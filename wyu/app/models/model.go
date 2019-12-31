package models

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/cast"
	"runtime"
	"wyu/app/exceptions"
	"wyu/configs"
)

const(
	SelectToOne string = "ONE"
	SelectToAll string = "ALL"
)

type Models struct {
	Engine *xorm.Engine
}

func New(engine *xorm.Engine) *Models {
	return &Models {
		Engine: engine,
	}
}

func (m *Models) Insert(xmSession *xorm.Session, data interface{}) (i int64, err error) {
	i, err = xmSession.Insert(data)
	return
}

func (m *Models) Update(xmSession *xorm.Session, dbInitialized configs.MdbInitialized, data interface{}) (i int64, err error) {
	if dbInitialized.Query == nil || dbInitialized.QueryArgs == nil {
		_, file, line, _ := runtime.Caller(1)
		err = exceptions.Err("m^aa", file, cast.ToString(line))
		return
	}

	i, err = xmSession.Where(dbInitialized.Query, dbInitialized.QueryArgs ...).Update(data)
	return
}

func (m *Models) Delete(xmSession *xorm.Session, dbInitialized configs.MdbInitialized, data interface{}) (i int64, err error) {
	if dbInitialized.Query == nil || dbInitialized.QueryArgs == nil {
		_, file, line, _ := runtime.Caller(1)
		err = exceptions.Err("m^aa", file, cast.ToString(line))
		return
	}

	i, err = xmSession.Where(dbInitialized.Query, dbInitialized.QueryArgs ...).Delete(data)
	return
}

func (m *Models) Total(data interface{}) (nums int64, err error) {
	engine := m.Engine.NewSession()
	defer engine.Close()

	nums, err = engine.Count(data)
	return
}

func (m *Models) Select(dbInitialized configs.MdbInitialized, data interface{}) (err error) {
	engine := m.Engine.NewSession()
	defer engine.Close()

	if dbInitialized.Columns != nil {
		engine = engine.Cols(dbInitialized.Columns ...)
	} else {
		engine = engine.Cols()
	}

	if dbInitialized.Query != nil && dbInitialized.QueryArgs != nil {
		engine = engine.Where(dbInitialized.Query, dbInitialized.QueryArgs ...)
	}

	if dbInitialized.Joins != nil {
		for _, join := range dbInitialized.Joins {
			engine = engine.Join(join[0].(string), join[1], join[2].(string))
		}
	}

	switch dbInitialized.Types {
	case SelectToOne:
		_, err = engine.Get(data)
		return

	case SelectToAll:
		err = engine.Find(data)
		return

	default:
		_, file, line, _ := runtime.Caller(1)
		err = exceptions.Err("m^ab", file, cast.ToString(line))
		return
	}
}
