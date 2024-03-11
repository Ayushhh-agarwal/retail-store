package mutex

import (
	"log"
	"net"
	"strconv"

	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func InitializeRedis(config RedisConfig) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort(config.Host, strconv.Itoa(int(config.Port))),
	})
	log.Println("redis initialise done")
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     "127.0.0.1",
		Port:     6379,
		Database: 0,
		Password: "",
	}
}

type RedisConfig struct {
	Host     string
	Port     int32
	Database int32
	Password string
}
