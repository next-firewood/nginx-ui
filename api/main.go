package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"rui/common/viper"
	_ "rui/docs"
	"rui/internal/repo"
	"rui/internal/router"
	"rui/internal/svc"
)

const _url = "./resource/config.yaml"

func main() {
	// 获取配置
	config := viper.InitConfig(_url)
	// 初始化服务上下文
	serviceContext := svc.NewServiceContext(config)
	// 初始化表
	if err := repo.InitTable(repo.GetDB(config.Database)); err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	// 初始化路由
	router.InitRouter(r, serviceContext)
	// 启动服务
	if err := r.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		log.Fatal(err)
	}
}
