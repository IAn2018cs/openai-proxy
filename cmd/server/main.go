package main

import (
	"log"

	"github.com/IAn2018cs/openai-proxy/config"
	"github.com/IAn2018cs/openai-proxy/internal/api"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 设置路由
	router := api.SetupRouter()

	// 启动服务器
	port := config.AppConfig.Server.Port
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
