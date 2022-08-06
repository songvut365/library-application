package database

import (
	"os"

	"github.com/go-redis/redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	HOST := os.Getenv("REDIS_ADDR")
	PORT := os.Getenv("REDIS_PORT")

	RDB = redis.NewClient(&redis.Options{
		Addr:     HOST + ":" + PORT,
		Password: "",
		DB:       0,
	})
}
