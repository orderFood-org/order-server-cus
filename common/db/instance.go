package db

import (
	"sync"

	"gorm.io/gorm"
)

var (
	dbInstance *Database
	once       sync.Once
)

// Database 数据库实例封装
type Database struct {
	db *gorm.DB
}

// GetInstance 获取数据库实例
func GetInstance() *Database {
	once.Do(func() {
		db, err := InitDB()
		if err != nil {
			panic(err)
		}
		dbInstance = &Database{db: db}
	})
	return dbInstance
}

// GetDB 获取GORM数据库实例
func (d *Database) GetDB() *gorm.DB {
	return d.db
}
