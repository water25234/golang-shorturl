package server

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/water25234/golang-shorturl/config"
	"github.com/water25234/golang-shorturl/core/log"
)

var client *redis.Client

func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.GetAppConfig().RedisHost + ":" + config.GetAppConfig().RedisPort,
		DB:   config.GetAppConfig().RedisDB,
	})

	client = redisClient

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
}

func SetRedis(key string, value string, num int) {
	log.Info("set redis")
	err := client.Set(key, value, time.Duration(num)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func GetRedis(key string) string {
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func GetKeys(key string) []string {
	val, err := client.Keys(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
