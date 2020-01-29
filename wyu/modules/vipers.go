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

func (cfg *vipers) Loading() *vipers {
	cfg.Env.SetConfigType(EnvType)
	cfg.Env.AddConfigPath(".")
	cfg.Env.SetConfigName(".env")

	if err := cfg.Env.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	cfg.Env.WatchConfig()
	cfg.Env.OnConfigChange(func (e fsnotify.Event){

	})

	return cfg
}

func (cfg *vipers) LoadInitializedFromYaml() *vipers {
	cfg.Env.AddConfigPath(EnvPath)
	cfg.Env.SetConfigType(EnvType)

	var env string = cfg.Env.GetString("env")
	if env == "" {
		panic("--env=? is not configured")
	}

	if ok, _ := UtilsStrContains(env, "dev","stg","prd"); ok == false {
		panic("--env=? must be in dev,stg,prd")
	}

	cfg.Env.SetConfigName(".env." + env)

	if err := cfg.Env.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	cfg.Env.WatchConfig()
	cfg.Env.OnConfigChange(func (e fsnotify.Event){})

	return cfg
}

func (cfg *vipers) GET(key string, val interface{}) interface{} {
	if cfg.Env.IsSet(key) {
		return cfg.Env.Get(key)
	}

	return val
}


