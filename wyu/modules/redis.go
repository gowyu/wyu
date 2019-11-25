package modules

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
)

func InstanceRedis() *rd {
	if len(RdEngines) == 0 {
		log.Fatal("Engine Redis Error, Plz Open the Redis Configure in .env.yaml")
	}

	if len(RdEngines) == 1 {
		return RdEngines[0]
	} else {
		return RdEngines[UtilsRandInt(0, len(RdEngines))]
	}
}

type rdSource struct {
	Addr string
	Password string
	DB int
	PoolSize int
}

type rd struct {
	r *redis.Client
	rdSource *rdSource
	mx sync.Mutex
}

var RdEngines []*rd

func NewRedis() *rd {
	return &rd{}
}

func (cache *rd) Engine() *redis.Client {
	if cache.r == nil {
		log.Println("Error Redis Engine")
		return nil
	}

	return cache.r
}

func (cache *rd) instance() *rd {
	cache.mx.Lock()
	defer cache.mx.Unlock()

	if cache.r != nil {
		return cache
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cache.rdSource.Addr,
		Password: cache.rdSource.Password,
		DB:       cache.rdSource.DB,
		PoolSize: cache.rdSource.PoolSize,
	})

	_, err := client.Ping().Result()
	if err != nil {
		if client != nil {
			client.Close()
		}

		panic(fmt.Sprintf("ping error[%s]\n", err.Error()))
	}

	cache.r = client
	return cache
}