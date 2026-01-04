package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func InitRedis() *redis.Client {

	host := viper.GetString("REDIS_HOST")
	port := viper.GetString("REDIS_PORT")

	fmt.Println(host, port)

	addr := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: viper.GetString("REDIS_PASSWORD"),
		DB: 0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect redis: %v ", err)
	}

	return client
}