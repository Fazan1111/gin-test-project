package redisLib

import (
	"fmt"
	envconfig "learnGin/src/common/envConfig"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func InitRedis() {
	redisHost := envconfig.GetEnv("REDIS_HOST")
	redisDB, _ := strconv.Atoi(envconfig.GetEnv("REDIS_DB"))
	client = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",      // no password set
		DB:       redisDB, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")

}

func GetRedisClient() *redis.Client {
	return client
}

func SetVal(key string, value string) {
	client.Set(key, value, time.Duration(time.Second*60))
}

func GetVal(key string) string {
	return client.Get(key).Val()
}
