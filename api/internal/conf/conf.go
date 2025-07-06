package conf

import (
	"rui/common/database"
	"rui/common/redis"
	"rui/internal/middleware"
)

type Conf struct {
	Mode     string
	Port     int
	Database *database.DBConf
	Auth     *middleware.AuthConf
	Cache    *redis.RdsConfig
}
