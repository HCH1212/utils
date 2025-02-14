package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// RedisClient 包含了 Redis 连接池和客户端
var RedisClient *redis.Client

// InitRedis 初始化 Redis 连接
func InitRedis(addr string, password string, db int) {
	// 创建 Redis 客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,            // Redis 地址
		Password:     password,        // Redis 密码
		DB:           db,              // Redis 数据库索引
		PoolSize:     10,              // 设置连接池大小
		MinIdleConns: 5,               // 设置最小空闲连接数
		IdleTimeout:  5 * time.Minute, // 设置空闲连接超时时间
	})

	// 测试 Redis 连接
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully")
}

// 获取 Redis 客户端
func GetRedisClient() *redis.Client {
	if RedisClient == nil {
		log.Fatal("Redis client is not initialized")
	}
	return RedisClient
}

// 关闭 Redis 连接
func CloseRedis() {
	if RedisClient != nil {
		err := RedisClient.Close()
		if err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		} else {
			fmt.Println("Redis connection closed")
		}
	}
}
