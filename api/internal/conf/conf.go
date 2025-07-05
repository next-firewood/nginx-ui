package conf

import "rui/internal/middleware"

type Conf struct {
	Mode string
	Port int
	Auth *middleware.AuthConf
}
