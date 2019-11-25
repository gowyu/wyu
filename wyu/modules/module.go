package modules

import (
	"github.com/spf13/pflag"
)

type wYu struct {
	Srcdb bool
	Redis bool
	I18nT bool
}

func init() {
	wYuEnv()

	if Env == nil {
		panic("Error Env in module.go - init func()")
	}

	var wYu *wYu
	UtilsMapToStruct(Env.GET("wYu", []interface{}{}), &wYu)

	if wYu.Srcdb {
		wYuSrcdb()
	}

	if wYu.Redis {
		wYuRedis()
	}

	if wYu.I18nT {
		wYuI18nT()
	}
}

func wYuEnv() {
	if Env != nil {
		return
	}

	pflag.String("env", "", "environment configure")
	pflag.Parse()

	var env Vipers = NewVipers()

	err := env.Vipers().Env.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err.Error())
	}

	Env = env.LoadInitializedFromYaml()
	if Env == nil {
		panic("Error Env In wYuEnv")
	}

	return
}

func wYuSrcdb() {
	var configs *dbConfigs

	envConfigure := Env.GET("DBClusters.Configure", map[string]interface{}{}).(map[string]interface{})
	if len(envConfigure) == 0 {
		configs = &dbConfigs{
			DriverName: "mysql",
			MaxOpen: 1000,
			MaxIdle: 500,
			ShowedSQL: false,
			CachedSQL: false,
		}
	} else {
		err := UtilsMapToStruct(envConfigure, &configs)
		if err != nil {
			panic("Error Env DBClusters.Configure Configures")
		}
	}

	envDatabases := Env.GET("DBClusters.Databases", map[string]interface{}{}).(map[string]interface{})
	if len(envDatabases) == 0 {
		panic("Error Env DBClusters.Databases Configures")
	}

	masterDB = make(map[string][]*db, 0)
	slaverDB = make(map[string][]*db, 0)

	for table, dbs := range envDatabases {
		for method, databases := range dbs.(map[string]interface{}) {
			var dbEngines []*db = make([]*db, len(databases.([]interface{})))
			for key, database := range databases.([]interface{}) {
				var cluster *dbCluster

				toMap := UtilsInterfaceToStringInMap(database.(map[interface{}]interface{}))
				toMap["Table"] = table

				err := UtilsMapToStruct(toMap, &cluster)
				if err != nil {
					panic(err.Error())
				}

				dbEngine := NewDB()
				dbEngine.dbCluster = cluster
				dbEngine.dbConfigs = configs

				dbEngines[key] = dbEngine.instance()
			}

			switch method {
				case "master":
					masterDB[table] = dbEngines
				case "slaver":
					slaverDB[table] = dbEngines
				default:
					continue
			}
		}
	}
}

func wYuRedis() {
	envRedis := Env.GET("Redis", []interface{}{}).([]interface{})
	if len(envRedis) < 1 {
		panic("redis configs error")
	}

	for _, rd := range envRedis {
		var src *rdSource
		toMap := UtilsInterfaceToStringInMap(rd.(map[interface{}]interface{}))

		err := UtilsMapToStruct(toMap, &src)
		if err != nil {
			continue
		}

		cache := NewRedis()
		cache.rdSource = src

		RdEngines = append(RdEngines, cache.instance())
	}
}

func wYuI18nT() {
	if err := NewI18N().Loading(); err != nil {
		panic(err.Error())
	}
}