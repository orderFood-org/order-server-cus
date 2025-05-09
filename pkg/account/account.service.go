package account

import (
	"orderFood-server-cus/common/db"
)

// Service 账户服务
type Service struct {
	db *db.Database
}

// NewService 创建账户服务
func NewService() *Service {
	return &Service{
		db: db.GetInstance(),
	}
}

// GetByID 根据ID获取账户
func (s *Service) GetByID(id uint) (*Account, error) {
	var account Account
	result := s.db.GetDB().First(&account, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

// GetByUsername 根据用户名获取账户
func (s *Service) GetByUsername(username string) (*Account, error) {
	var account Account
	result := s.db.GetDB().Where("username = ?", username).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

// GetByEmail 根据邮箱获取账户
func (s *Service) GetByEmail(email string) (*Account, error) {
	var account Account
	result := s.db.GetDB().Where("email = ?", email).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

// Create 创建账户
func (s *Service) Create(account *Account) error {
	return s.db.GetDB().Create(account).Error
}

// Update 更新账户
func (s *Service) Update(account *Account) error {
	return s.db.GetDB().Save(account).Error
}

// Delete 删除账户
func (s *Service) Delete(id uint) error {
	return s.db.GetDB().Delete(&Account{}, id).Error
}

// List 获取账户列表
func (s *Service) List(page, pageSize int) ([]Account, int64, error) {
	var accounts []Account
	var total int64

	query := s.db.GetDB().Model(&Account{})
	query.Count(&total)

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	result := query.Find(&accounts)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return accounts, total, nil
}
