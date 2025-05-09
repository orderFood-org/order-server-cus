package dish

import (
	"orderFood-server-cus/common/db"
)

// Service 菜品服务
type Service struct {
	db *db.Database
}

// NewService 创建菜品服务
func NewService() *Service {
	return &Service{
		db: db.GetInstance(),
	}
}

// GetByID 根据ID获取菜品
func (s *Service) GetByID(id uint) (*Dish, error) {
	var dish Dish
	result := s.db.GetDB().Preload("Category").First(&dish, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dish, nil
}

// GetByCategoryID 根据分类ID获取菜品列表
func (s *Service) GetByCategoryID(categoryID uint) ([]Dish, error) {
	var dishes []Dish
	result := s.db.GetDB().Where("category_id = ?", categoryID).
		Preload("Category").
		Order("sort_order ASC").
		Find(&dishes)
	if result.Error != nil {
		return nil, result.Error
	}
	return dishes, nil
}

// Create 创建菜品
func (s *Service) Create(dish *Dish) error {
	return s.db.GetDB().Create(dish).Error
}

// Update 更新菜品
func (s *Service) Update(dish *Dish) error {
	return s.db.GetDB().Save(dish).Error
}

// Delete 删除菜品
func (s *Service) Delete(id uint) error {
	return s.db.GetDB().Delete(&Dish{}, id).Error
}

// List 获取菜品列表
func (s *Service) List(page, pageSize int) ([]Dish, int64, error) {
	var dishes []Dish
	var total int64

	query := s.db.GetDB().Model(&Dish{}).Preload("Category")
	query.Count(&total)

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	result := query.Order("sort_order ASC").Find(&dishes)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return dishes, total, nil
}

// GetSpecials 获取特色菜品
func (s *Service) GetSpecials() ([]Dish, error) {
	var dishes []Dish
	result := s.db.GetDB().Where("is_special = ?", true).
		Preload("Category").
		Order("sort_order ASC").
		Find(&dishes)
	if result.Error != nil {
		return nil, result.Error
	}
	return dishes, nil
}
