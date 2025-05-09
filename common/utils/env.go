package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv 加载环境变量
func LoadEnv() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	// 尝试加载对应环境的配置文件
	err := godotenv.Load(".env." + env)
	if err != nil {
		// 如果环境特定配置不存在，加载默认配置
		err = godotenv.Load()
		if err != nil {
			// 仅记录日志，不终止程序
			log.Printf("Warning: No .env file found. Using environment variables")
		}
	}
}

// GetEnv 从环境变量获取配置，如果不存在则使用默认值
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
