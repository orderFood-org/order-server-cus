package service

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewService() *Service {
	s := &Service{}

	connectDB(s)
	connectRDB(s)

	return s
}

func connectDB(s *Service) {
	mysqlDB, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// mysqlDB.AutoMigrate(&model.Account{})
	mysqlDB = mysqlDB.Debug() //

	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	s.db = mysqlDB
}

func connectRDB(s *Service) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	s.rdb = rdb
}
