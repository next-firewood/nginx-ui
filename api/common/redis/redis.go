package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RdsConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func (r *RdsConfig) NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
	})

	return client
}
