#!/bin/bash

echo "启动点餐系统服务端..."

# 清理旧的构建缓存
go clean -cache

# 获取依赖
go mod tidy

# 运行应用
go run cmd/main.go 