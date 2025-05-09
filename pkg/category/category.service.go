package category

import (
	"orderFood-server-cus/common/db"
)

// Service 分类服务
type Service struct {
	db *db.Database
}

// NewService 创建分类服务
func NewService() *Service {
	return &Service{
		db: db.GetInstance(),
	}
}

// GetByID 根据ID获取分类
func (s *Service) GetByID(id uint) (*Category, error) {
	var category Category
	result := s.db.GetDB().First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

// Create 创建分类
func (s *Service) Create(category *Category) error {
	return s.db.GetDB().Create(category).Error
}

// Update 更新分类
func (s *Service) Update(category *Category) error {
	return s.db.GetDB().Save(category).Error
}

// Delete 删除分类
func (s *Service) Delete(id uint) error {
	return s.db.GetDB().Delete(&Category{}, id).Error
}

// List 获取分类列表
func (s *Service) List() ([]Category, error) {
	var categories []Category
	result := s.db.GetDB().Order("sort_order ASC").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// GetActive 获取活跃分类列表
func (s *Service) GetActive() ([]Category, error) {
	var categories []Category
	result := s.db.GetDB().Where("is_active = ?", true).
		Order("sort_order ASC").
		Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
