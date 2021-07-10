package cache

import (
	"awesome-runner/src/config"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func initRedisClient() *redis.Client {
	options, _ := redis.ParseURL(fmt.Sprintf(`redis://%s/%d`, config.Cnf.RedisDNS, config.Cnf.DefaultDb))
	options.PoolSize = 20
	options.MinIdleConns = 10

	return redis.NewClient(options)
}

// get *redis.Client
func GetReidsInstance() *redis.Client {
	if RedisClient == nil {
		var once sync.Once
		once.Do(func() {
			RedisClient = initRedisClient()
		})
	}
	return RedisClient
}
