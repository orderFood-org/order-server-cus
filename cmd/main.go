package main

import (
	"fmt"
	"log"
	"orderFood-server-cus/common/db"
	"orderFood-server-cus/common/middleware"
	"orderFood-server-cus/common/utils"
	"orderFood-server-cus/pkg/account"
	"orderFood-server-cus/pkg/category"
	"orderFood-server-cus/pkg/dish"
	"orderFood-server-cus/pkg/order"
	"orderFood-server-cus/pkg/token"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载环境变量
	utils.LoadEnv()

	// 初始化数据库
	dbInstance := db.GetInstance()
	if dbInstance == nil {
		log.Fatalf("Failed to initialize database")
	}

	// 注册模型到迁移列表
	db.RegisterModel(&account.Account{})
	db.RegisterModel(&token.Token{})
	db.RegisterModel(&category.Category{})
	db.RegisterModel(&dish.Dish{})
	db.RegisterModel(&order.Order{})

	// 执行数据库迁移
	if err := db.MigrateDB(dbInstance.GetDB()); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化Gin框架
	r := gin.Default()

	// CORS配置
	r.Use(middleware.CORS())

	// API路由组
	api := r.Group("/api")
	{
		// 注册各模块路由
		token.RegisterRoutes(api)
		account.RegisterRoutes(api)
		dish.RegisterRoutes(api)
		category.RegisterRoutes(api)
		order.RegisterRoutes(api)
	}

	// 获取端口
	port := utils.GetEnv("PORT", "3003")

	// 启动服务器
	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
