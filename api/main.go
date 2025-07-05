package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rui/common/viper"
	_ "rui/docs"
	"rui/internal/router"
	"rui/internal/svc"
)

const _url = "./resource/config.yaml"

func main() {
	config := viper.InitConfig(_url)

	r := gin.New()

	serviceContext := svc.NewServiceContext(config)

	router.InitRouter(r, serviceContext)

	if config.Mode == "dev" {
		// Swagger 接口
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run(fmt.Sprintf(":%d", config.Port))
}
