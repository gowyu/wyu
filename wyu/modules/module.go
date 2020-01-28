package modules

import (
	"fmt"
	"github.com/spf13/pflag"
	"runtime"
)

var (
	YuEnv *vipers
	YuToken *Token
	YuEmail *Mail
)

type Yu struct {
	Srcdb bool
	Redis bool
	I18nT bool
	Email bool
}

func init() {
	modules := New()
	modules.wYuEnv()
	modules.wYuToken()

	var Yu *Yu
	UtilsMapToStruct(Env.GET("Yu", []interface{}{}), &Yu)

	if Yu.Srcdb {
		modules.wYuSrcdb()
	}

	if Yu.Redis {
		modules.wYuRedis()
	}

	if Yu.I18nT {
		modules.wYuI18nT()
	}

	if Yu.Email {
		modules.wYuEmail()
	}
}

type modules struct {

}

func New() *modules {
	return &modules {}
}

func (module *modules) wYuEnv() *modules {
	if Env != nil {
		return module
	}

	pflag.String("env", "", "environment configure")
	pflag.Parse()

	env := NewVipers()
	err := env.Env.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err.Error())
	}

	Env = env.LoadInitializedFromYaml()
	if Env == nil {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("Error Env In wYuEnv » %v » %v", file, line))
	}

	return module
}

func (module *modules) wYuToken() *modules {
	YuToken = NewToken()
	return module
}

func (module *modules) wYuSrcdb() *modules {
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
			_, file, line, _ := runtime.Caller(1)
			panic(fmt.Sprintf("Error Env DBClusters.Configure Configures » %v » %v", file, line))
		}
	}

	envDatabases := Env.GET("DBClusters.Databases", map[string]interface{}{}).(map[string]interface{})
	if len(envDatabases) == 0 {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("Error Env DBClusters.Databases Configures » %v » %v", file, line))
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

	return module
}

func (module *modules) wYuRedis() *modules {
	env := Env.GET("Redis", []interface{}{}).([]interface{})
	if len(env) < 1 {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("redis configs error in modules/module.go » %v » %v", file, line))
	}

	for _, rd := range env {
		var src *rdSource
		toMap := UtilsInterfaceToStringInMap(rd.(map[interface{}]interface{}))

		if err := UtilsMapToStruct(toMap, &src); err != nil {
			continue
		}

		cache := NewRedis()
		cache.rdSource = src

		RdEngines = append(RdEngines, cache.instance())
	}

	return module
}

func (module *modules) wYuI18nT() *modules {
	if err := NewI18N().Loading(); err != nil {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintf("%v | %v | %v", err.Error(), file, line))
	}

	return module
}

func (module *modules) wYuEmail() *modules {
	var cfg *MailConfigs
	env := Env.GET("YuMail", map[string]interface{}{}).(map[string]interface{})
	UtilsMapToStruct(env, &cfg)

	YuEmail = NewMail(cfg)
	return module
}