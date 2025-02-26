package handler

import (
	"net/http"

	"github.com/IAn2018cs/openai-proxy/internal/model"
	"github.com/IAn2018cs/openai-proxy/internal/service"
	"github.com/gin-gonic/gin"
)

type OpenAIHandler struct {
	openaiService *service.OpenAIService
}

func NewOpenAIHandler(openaiService *service.OpenAIService) *OpenAIHandler {
	return &OpenAIHandler{
		openaiService: openaiService,
	}
}

// HandleChatCompletion 处理聊天完成请求
func (h *OpenAIHandler) HandleChatCompletion(c *gin.Context) {
	var req model.OpenAIRequest

	// 解析请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查是否有自定义的base_url
	baseURL := c.GetHeader("X-OpenAI-BaseURL")
	if baseURL != "" {
		h.openaiService.SetBaseURL(baseURL)
	}

	// 检查是否有自定义的API密钥
	apiKey := c.GetHeader("X-OpenAI-APIKey")
	if apiKey != "" {
		h.openaiService.SetAPIKey(apiKey)
	}

	// 调用服务
	resp, err := h.openaiService.CreateChatCompletion(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, resp)
}
