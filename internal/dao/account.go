package dao

import (
	"fmt"
	"orderFood-server-cus/internal/model"

	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

func (a *AccountDao) IsExist(username string) (bool, error) {
	account := &model.Account{}

	err := a.db.Where("username = ?", username).First(account).Error
	if err == gorm.ErrRecordNotFound { // 符合预期的错误
		return false, err
	}
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false, err
	}

	return true, nil
}

func (a *AccountDao) Create(account *model.Account) error {
	fmt.Print("account: ", account)
	return a.db.Create(&account).Error
}

func (a *AccountDao) GetAccountByUsername(username string) (*model.Account, error) {
	account := &model.Account{}
	err := a.db.Where("username = ?", username).First(account).Error
	if err != nil {
		fmt.Printf("login Error: %v", err)
		return nil, err
	}

	return account, nil
}
