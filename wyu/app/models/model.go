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
	SelectToASC string = "ASC"
	SelectToESC string = "DESC"
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

	/**
	 * Todo: Select Table & Columns
	 */
	if dbInitialized.Table != "" && dbInitialized.Field != "" {
		if ok, _ := engine.IsTableExist(dbInitialized.Table); ok == false {
			_, file, line, _ := runtime.Caller(1)
			err = exceptions.Err("m^ad", file, cast.ToString(line))
			return
		}
		
		engine = engine.Table(dbInitialized.Table).Select(dbInitialized.Field)
	} else {
		if dbInitialized.Columns != nil {
			engine = engine.Cols(dbInitialized.Columns ...)
		} else {
			engine = engine.Cols()
		}
	}

	/**
	 * Todo: Add Join Table (INNER|LEFT|RIGHT)
	 */
	if dbInitialized.Joins != nil {
		for _, join := range dbInitialized.Joins {
			if len(join) > 2 {
				engine = engine.Join(join[0].(string), join[1], join[2].(string))
			}
		}
	}

	/**
	 * Todo: Add Condition
	 */
	if dbInitialized.Query != nil && dbInitialized.QueryArgs != nil {
		engine = engine.Where(dbInitialized.Query, dbInitialized.QueryArgs ...)
	}

	switch dbInitialized.Types {
	case SelectToOne:
		_, err = engine.Get(data)
		return

	case SelectToAll:
		/**
		 * Todo: Add OrderBy (ASC|DESC)
		 */
		if len(dbInitialized.OrderArgs) > 0 {
			if dbInitialized.OrderType == SelectToASC {
				engine = engine.Asc(dbInitialized.OrderArgs ...)
			}

			if dbInitialized.OrderType == SelectToESC {
				engine = engine.Desc(dbInitialized.OrderArgs ...)
			}
		}

		/**
		 * Todo: Limit & Start
		 */
		if dbInitialized.Limit != 0 && len(dbInitialized.Start) > 0 {
			engine = engine.Limit(dbInitialized.Limit, dbInitialized.Start ...)
		}

		err = engine.Find(data)
		return

	default:
		_, file, line, _ := runtime.Caller(1)
		err = exceptions.Err("m^ab", file, cast.ToString(line))
		return
	}
}

func (m *Models) ToSerialized() {

}
