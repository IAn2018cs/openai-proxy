# OpenAI API Proxy

一个简单的Go后端项目，用于转发OpenAI的chat completion接口请求。

## 功能

- 转发非流式OpenAI chat completion接口请求
- 支持自定义模型和system消息
- 支持自定义OpenAI基础URL
- 支持API密钥认证

## 配置

在`config/config.yaml`中设置配置：

```yaml
server:
  port: "8080"

openai:
  base_url: "https://api.openai.com"
  api_key: "your-openai-api-key"

auth:
  api_key: "your-auth-api-key"
```

## API使用

### 请求格式
```
POST /v1/chat
Authorization: Bearer your-auth-api-key
Content-Type: application/json

{
  "model": "gpt-4o-mini",
  "system": "你是一个有用的助手",
  "content": "你好，请介绍一下自己",
  "options": {
    "temperature": 0.7,
    "max_tokens": 1000
  }
}
```

### 自定义 OpenAI 基础 URL
可以通过设置X-OpenAI-BaseURL头来自定义OpenAI的基础URL：
```
X-OpenAI-BaseURL: https://your-custom-openai-endpoint.com
```

### 自定义 OpenAI API 密钥
可以通过设置X-OpenAI-APIKey头来自定义OpenAI的API密钥：
```
X-OpenAI-APIKey: your-custom-openai-api-key
```

## 运行
```bash
go run cmd/server/main.go
```

## Docker 部署

### 使用 Docker 构建和运行

```bash
# 构建 Docker 镜像
docker build -t openai-proxy .

# 运行容器
docker run -d -p 8080:8080 -v $(pwd)/config:/root/config --name openai-proxy openai-proxy
```

### 使用 Docker Compose 部署
```bash
# 启动服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

### 配置说明

在使用 Docker 部署时，可以通过挂载配置文件或设置环境变量来配置应用：

1. 挂载配置文件：
```bash
docker run -d -p 8080:8080 -v $(pwd)/config:/root/config --name openai-proxy openai-proxy
```
2. 使用环境变量（优先级高于配置文件）：
```bash
docker run -d -p 8080:8080 \
  -e OPENAI_BASE_URL=https://api.openai.com \
  -e OPENAI_API_KEY=your-openai-api-key \
  -e AUTH_API_KEY=your-auth-api-key \
  -e SERVER_PORT=8080 \
  --name openai-proxy openai-proxy
```
