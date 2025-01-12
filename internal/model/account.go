package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	NickName string `gorm:"column:nick_name" json:"nick_name"`
}

func (a *Account) TableName() string {
	table := "order_food.account"
	return table
}
