package modules

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"strings"
	"sync"
	"time"
)

func InstanceClusterDB(db string, selector ...int) *db {
	if len(masterDB[strings.ToLower(db)]) == 0 {
		log.Println("MasterDB Error")
		return nil
	}

	if len(slaverDB[strings.ToLower(db)]) == 0 {
		log.Println("SlaverDB Error")
		return nil
	}

	if len(selector) > 0 {
		return masterDB[strings.ToLower(db)][UtilsRandInt(0, len(masterDB[strings.ToLower(db)]))]
	} else {
		return slaverDB[strings.ToLower(db)][UtilsRandInt(0, len(slaverDB[strings.ToLower(db)]))]
	}
}

var (
	SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

	masterDB map[string][]*db
	slaverDB map[string][]*db
)

type dbCluster struct {
	Host      	string
	Port      	int
	Table     	string
	Username  	string
	Password  	string
}

type dbConfigs struct {
	DriverName  string
	MaxOpen   	int
	MaxIdle   	int
	ShowedSQL 	bool
	CachedSQL 	bool
}

type db struct {
	engine *xorm.Engine
	mx sync.Mutex
	dbCluster *dbCluster
	dbConfigs *dbConfigs
}

func NewDB() *db {
	return &db{}
}

func (odbc *db) Engine() *xorm.Engine {
	if odbc.engine == nil {
		log.Println("Error DB Engine")
		return nil
	}

	return odbc.engine
}

func (odbc *db) instance() *db {
	odbc.mx.Lock()
	defer odbc.mx.Unlock()

	if odbc.engine != nil {
		return odbc
	}

	if odbc.dbConfigs == nil || odbc.dbCluster == nil {
		log.Fatal("Error DB Data Source Cluster or Configs")
	}

	driverFormat := "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	driverSource := fmt.Sprintf(
		driverFormat,
		odbc.dbCluster.Username,
		odbc.dbCluster.Password,
		odbc.dbCluster.Host,
		odbc.dbCluster.Port,
		odbc.dbCluster.Table,
	)

	engine, err := xorm.NewEngine(odbc.dbConfigs.DriverName, driverSource)
	if err != nil {
		if engine != nil {
			engine.Close()
		}

		log.Fatalf("db.DbInstanceMaster, %s", err.Error())
	}

	engine.SetMaxOpenConns(odbc.dbConfigs.MaxOpen)
	engine.SetMaxIdleConns(odbc.dbConfigs.MaxIdle)

	engine.ShowSQL(odbc.dbConfigs.ShowedSQL)
	engine.SetTZDatabase(SysTimeLocation)

	if odbc.dbConfigs.CachedSQL {
		cached := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cached)
	}

	odbc.engine = engine
	return odbc
}