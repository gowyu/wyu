package modules

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	EnvPath string = "."
	EnvType string = "yaml"
)

type vipers struct {
	Env *viper.Viper
}

var Env *vipers

func NewVipers() *vipers {
	return &vipers{
		Env:viper.New(),
	}
}

func (conf *vipers) Loading() *vipers {
	conf.Env.SetConfigType(EnvType)
	conf.Env.AddConfigPath(".")
	conf.Env.SetConfigName(".env")

	if err := conf.Env.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	conf.Env.WatchConfig()
	conf.Env.OnConfigChange(func (e fsnotify.Event){

	})

	return conf
}

func (conf *vipers) LoadInitializedFromYaml() *vipers {
	conf.Env.AddConfigPath(EnvPath)
	conf.Env.SetConfigType(EnvType)

	var env string = conf.Env.GetString("env")
	if env == "" {
		panic("--env=? is not configured")
	}

	if ok, _ := UtilsStrContains(env, "dev","stg","prd"); ok == false {
		panic("--env=? must be in dev,stg,prd")
	}

	conf.Env.SetConfigName(".env." + env)

	if err := conf.Env.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	conf.Env.WatchConfig()
	conf.Env.OnConfigChange(func (e fsnotify.Event){})

	return conf
}

func (conf *vipers) GET(key string, val interface{}) (res interface{}) {
	res = val
	if conf.Env.IsSet(key) {
		res = conf.Env.Get(key)
	}

	return
}


