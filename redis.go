package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
)

func CreateNewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:   0,
	})

	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	return rdb
}
