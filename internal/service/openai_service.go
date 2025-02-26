package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/IAn2018cs/openai-proxy/config"
	"github.com/IAn2018cs/openai-proxy/internal/model"
)

type OpenAIService struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

func NewOpenAIService() *OpenAIService {
	return &OpenAIService{
		client:  &http.Client{},
		baseURL: config.AppConfig.OpenAI.BaseURL,
		apiKey:  config.AppConfig.OpenAI.APIKey,
	}
}

// SetBaseURL 允许自定义基础URL
func (s *OpenAIService) SetBaseURL(baseURL string) {
	if baseURL != "" {
		s.baseURL = baseURL
	}
}

// SetAPIKey 设置API密钥
func (s *OpenAIService) SetAPIKey(apiKey string) {
	if apiKey != "" {
		s.apiKey = apiKey
	}
}

// CreateChatCompletion 转发请求到OpenAI
func (s *OpenAIService) CreateChatCompletion(req *model.OpenAIRequest) (*model.OpenAIResponse, error) {
	// 构建OpenAI请求
	messages := []model.Message{
		{
			Role:    "system",
			Content: req.System,
		},
		{
			Role:    "user",
			Content: req.Content,
		},
	}

	chatReq := model.ChatCompletionRequest{
		Model:    req.Model,
		Messages: messages,
		Options:  req.Options,
	}

	jsonData, err := json.Marshal(chatReq)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	// 构建URL
	url := path.Join(s.baseURL, "v1/chat/completions")
	if s.baseURL[len(s.baseURL)-1] != '/' {
		url = s.baseURL + "/v1/chat/completions"
	}

	// 创建HTTP请求
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+s.apiKey)

	// 发送请求
	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request to OpenAI: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenAI API error: %s, status code: %d", string(body), resp.StatusCode)
	}

	// 解析响应
	var openAIResp model.OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &openAIResp, nil
}
