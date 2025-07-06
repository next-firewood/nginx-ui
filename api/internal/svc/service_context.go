package svc

import (
	"github.com/redis/go-redis/v9"
	"rui/internal/conf"
	"rui/internal/repo"
)

type ServiceContext struct {
	Config *conf.Conf
	Redis  *redis.Client
	Repo   repo.GlobalRepo
}

func NewServiceContext(c *conf.Conf) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  c.Cache.NewRedisClient(),
		Repo:   repo.NewGlobalDb(repo.GetDB(c.Database)),
	}
}
