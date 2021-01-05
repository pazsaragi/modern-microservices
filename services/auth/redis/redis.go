package redis

import (
	"os"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
