package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	OpenAI struct {
		BaseURL string `mapstructure:"base_url"`
		APIKey  string `mapstructure:"api_key"`
	} `mapstructure:"openai"`
	Auth struct {
		APIKey string `mapstructure:"api_key"`
	} `mapstructure:"auth"`
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("openai.base_url", "https://api.openai.com")

	// 支持环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Config file not found or invalid: %v", err)
		log.Println("Using environment variables or default values")
	}

	// 绑定环境变量
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("openai.base_url", "OPENAI_BASE_URL")
	viper.BindEnv("openai.api_key", "OPENAI_API_KEY")
	viper.BindEnv("auth.api_key", "AUTH_API_KEY")

	// 解析配置到结构体
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	// 验证必要的配置
	if AppConfig.OpenAI.APIKey == "" {
		log.Println("Warning: OpenAI API Key not set")
	}
	
	if AppConfig.Auth.APIKey == "" {
		log.Println("Warning: Auth API Key not set")
	}
}