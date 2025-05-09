package db

import (
	"fmt"
	"log"
	"orderFood-server-cus/common/utils"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL 驱动
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	instance *gorm.DB
)

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return instance
}

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	if instance != nil {
		return instance, nil
	}

	// 数据库配置
	dbUser := utils.GetEnv("DB_USER", "postgres")
	dbPassword := utils.GetEnv("DB_PASSWORD", "123456")
	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbPort := utils.GetEnv("DB_PORT", "5432")
	dbName := utils.GetEnv("DB_NAME", "order_food")
	sslMode := utils.GetEnv("DB_SSLMODE", "disable")

	// 构建DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)

	// 日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接数据库
	var err error
	instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // PostgreSQL 无需表前缀
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		log.Printf("无法连接到数据库 %s: %v，尝试创建数据库...", dbName, err)

		// 尝试连接到 postgres 数据库并创建目标数据库
		pgDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
			dbHost, dbPort, dbUser, dbPassword, sslMode)

		instance, err = gorm.Open(postgres.Open(pgDSN), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("无法连接到 postgres 数据库: %v", err)
		}

		// 创建数据库
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if err := instance.Exec(createDBSQL).Error; err != nil {
			return nil, fmt.Errorf("创建数据库失败: %v", err)
		}

		log.Printf("成功创建数据库: %s", dbName)

		// 重新连接到新创建的数据库
		instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
		})

		if err != nil {
			return nil, fmt.Errorf("无法连接到新创建的数据库: %v", err)
		}
	}

	// 获取底层SQL连接池并配置
	sqlDB, err := instance.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return instance, nil
}
