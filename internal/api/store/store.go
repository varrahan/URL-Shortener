package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/varrahan/url-shortener/internal/api/utils"
	"log"
	"time"
)

type Storage struct {
	redisClient *redis.Client
}

var (
	store = &Storage{}
	ctx = context.Background()
)

const CacheDuration = time.Hour

func InitStore() *Storage {

	redisUrl := utils.GetEnv("REDIS_ADDR", "localhost:6379")
	var redisClient *redis.Client

	if redisUrl == "localhost:6379" {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisUrl,
			Password: "",
			DB:       0,
		})
	} else {
		opt, err := redis.ParseURL(redisUrl)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to redis. Error: %v", err))
		}
		redisClient = redis.NewClient(opt)
	}

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error with redis initialization: %v\n", err))
	}

	log.Printf("Redis successfully initialized: pong message = %s\n", pong)
	store.redisClient = redisClient
	return store
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := store.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}


func RetrieveInitialUrl(shortUrl string) string {
	result, err := store.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}