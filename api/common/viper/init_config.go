package viper

import (
	"github.com/spf13/viper"
	"log"
	"rui/internal/conf"
)

// InitConfig 初始化配置文件
func InitConfig(url string) *conf.Conf {
	v := viper.New()
	// 1. 读取配置文件
	v.SetConfigFile(url)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	config := &conf.Conf{}
	if err := v.Unmarshal(config); err != nil {
		log.Fatal(err)
	}

	return config
}
