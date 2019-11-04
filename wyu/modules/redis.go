package modules

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
)

func InstanceRedis() *rd {
	if len(RdEngines) < 1 {
		log.Fatal("Engine Redis Error, Plz Open the Redis Configure in .env.yaml")
	}

	return RdEngines[UtilsRandInt(0, len(RdEngines))]
}

type rdSource struct {
	Addr string
	Password string
	DB int
	PoolSize int
}

type Rd interface {
	Rd() *rd
	Engine() *redis.Client
	Instance() *rd
}

type rd struct {
	r *redis.Client
	mx sync.Mutex
	RdSource *rdSource
}

var (
	_ Rd = &rd{}
	RdEngines []*rd
)

func NewRedis() *rd {
	return &rd{}
}

func (cache *rd) Rd() *rd {
	return cache
}

func (cache *rd) Engine() *redis.Client {
	if cache.r == nil {
		log.Println("Error Redis Engine")
		return nil
	}

	return cache.r
}

func (cache *rd) Instance() *rd {
	cache.mx.Lock()
	defer cache.mx.Unlock()

	if cache.r != nil {
		return cache
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cache.RdSource.Addr,
		Password: cache.RdSource.Password,
		DB:       cache.RdSource.DB,
		PoolSize: cache.RdSource.PoolSize,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	cache.r = client
	return cache
}


