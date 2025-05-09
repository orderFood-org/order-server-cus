package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// 模型接口，所有需要自动迁移的模型都需要实现这个接口
type Model interface {
	TableName() string
}

// RegisterModels 注册所有需要迁移的模型
var models []interface{}

// RegisterModel 注册模型到迁移列表
func RegisterModel(model interface{}) {
	models = append(models, model)
}

// MigrateDB 执行数据库迁移
func MigrateDB(db *gorm.DB) error {
	log.Println("开始自动迁移数据库模型...")

	// 确保使用public模式
	db.Exec("SET search_path TO public")

	// 尝试直接删除表并重新创建 - 仅在开发环境使用
	// 对于生产环境应该使用更安全的迁移方式
	tableNames := []string{"users", "tokens", "categories", "dishes", "orders"}
	for _, tableName := range tableNames {
		dropSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
		if err := db.Exec(dropSQL).Error; err != nil {
			log.Printf("删除表 %s 失败: %v", tableName, err)
			// 继续尝试其他表
		} else {
			log.Printf("已删除表 %s 准备重新创建", tableName)
		}
	}

	// 创建新的配置会话
	newSession := db.Session(&gorm.Session{
		SkipDefaultTransaction: true, // 跳过默认事务
		AllowGlobalUpdate:      true, // 允许全局更新
	})

	// 迁移所有已注册的模型
	err := newSession.AutoMigrate(models...)
	if err != nil {
		log.Printf("模型自动迁移失败: %v", err)
		return err
	}

	log.Println("数据库模型迁移完成")
	return nil
}
