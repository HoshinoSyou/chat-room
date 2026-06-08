package dao

import (
	"chat-room/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func InitRedis() error {
	redisConfig := config.AppConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	result, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("无法连接到 Redis 服务器，错误信息：%v", err)
		return err
	}
	log.Printf("Redis 连接成功！%v", result)
	Rdb = client
	return nil
}
