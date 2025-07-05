package svc

import "rui/internal/conf"

type ServiceContext struct {
	Config *conf.Conf
}

func NewServiceContext(c *conf.Conf) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
