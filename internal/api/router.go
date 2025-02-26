package api

import (
	"net/http"

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

	// 健康检查端点 - 不需要认证
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API 路由组 - 需要认证
	api := router.Group("/")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/v1/chat", openaiHandler.HandleChatCompletion)
	}

	return router
}
