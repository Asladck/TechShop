package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	logrus.Println("Redis Connected")
	return client
}
