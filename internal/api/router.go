package api

import (
	"github.com/IAn2018cs/openai-proxy/internal/api/handler"
	"github.com/IAn2018cs/openai-proxy/internal/api/middleware"
	"github.com/IAn2018cs/openai-proxy/internal/service"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 创建服务和处理器
	openaiService := service.NewOpenAIService()
	openaiHandler := handler.NewOpenAIHandler(openaiService)

	// 添加中间件
	router.Use(middleware.AuthMiddleware())

	// 注册路由
	router.POST("/v1/chat", openaiHandler.HandleChatCompletion)

	return router
}
