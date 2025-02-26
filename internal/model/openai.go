package model

// OpenAIRequest 定义了接收的请求结构
type OpenAIRequest struct {
	Model   string         `json:"model" binding:"required"`
	System  string         `json:"system"`
	Content string         `json:"content" binding:"required"`
	Options RequestOptions `json:"options,omitempty"`
}

// RequestOptions 定义了请求选项
type RequestOptions struct {
	Temperature         float64 `json:"temperature,omitempty"`
	MaxTokens           int     `json:"max_tokens,omitempty"`
	MaxCompletionTokens int     `json:"max_completion_tokens,omitempty"`
	ReasoningEffort     float64 `json:"reasoning_effort,omitempty"`
}

// ChatCompletionRequest 定义了发送给OpenAI的请求结构
type ChatCompletionRequest struct {
	Model               string    `json:"model"`
	Messages            []Message `json:"messages"`
	Temperature         float64   `json:"temperature,omitempty"`
	MaxTokens           int       `json:"max_tokens,omitempty"`
	MaxCompletionTokens int       `json:"max_completion_tokens,omitempty"`
	ReasoningEffort     float64   `json:"reasoning_effort,omitempty"`
}

// Message 定义了消息结构
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse 定义了OpenAI的响应结构
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice 定义了OpenAI响应中的选择
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage 定义了OpenAI响应中的用量信息
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// SimpleResponse 定义了简化后的响应结构，只包含消息内容
type SimpleResponse struct {
	Content string `json:"content"`
}
