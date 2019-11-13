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
	if len(masterDB[strings.ToLower(db)]) < 1 || len(slaverDB[strings.ToLower(db)]) < 1 {
		log.Fatal("MasterDB or SlaverDB Error")
	}

	if len(selector) > 0 {
		return masterDB[strings.ToLower(db)][UtilsRandInt(0, len(masterDB[strings.ToLower(db)]))]
	} else {
		return slaverDB[strings.ToLower(db)][UtilsRandInt(0, len(slaverDB[strings.ToLower(db)]))]
	}
}

var (
	_ DB = &db{}
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

type DB interface {
	DB() *db
	Instance() *db
	Engine() *xorm.Engine
}

type db struct {
	engine *xorm.Engine
	mx sync.Mutex
	DBCluster *dbCluster
	DBConfigs *dbConfigs
}

func NewDB() *db {
	return &db{}
}

func (odbc *db) DB() *db {
	return odbc
}

func (odbc *db) Engine() *xorm.Engine {
	if odbc.engine == nil {
		log.Println("Error DB Engine")
		return nil
	}

	return odbc.engine
}

func (odbc *db) Instance() *db {
	odbc.mx.Lock()
	defer odbc.mx.Unlock()

	if odbc.engine != nil {
		return odbc
	}

	if odbc.DBConfigs == nil || odbc.DBCluster == nil {
		log.Println("Error DB Data Source Cluster or Configs")
	}

	driverFormat := "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	driverSource := fmt.Sprintf(
		driverFormat,
		odbc.DBCluster.Username,
		odbc.DBCluster.Password,
		odbc.DBCluster.Host,
		odbc.DBCluster.Port,
		odbc.DBCluster.Table,
	)

	engine, err := xorm.NewEngine(odbc.DBConfigs.DriverName, driverSource)
	if err != nil {
		log.Fatalf("db.DbInstanceMaster, %s", err.Error())
	}

	engine.SetMaxOpenConns(odbc.DBConfigs.MaxOpen)
	engine.SetMaxIdleConns(odbc.DBConfigs.MaxIdle)

	engine.ShowSQL(odbc.DBConfigs.ShowedSQL)
	engine.SetTZDatabase(SysTimeLocation)

	if odbc.DBConfigs.CachedSQL {
		cached := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cached)
	}

	odbc.engine = engine
	return odbc
}