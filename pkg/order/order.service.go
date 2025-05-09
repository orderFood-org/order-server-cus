package order

import (
	"orderFood-server-cus/common/db"
)

// Service 订单服务
type Service struct {
	db *db.Database
}

// NewService 创建订单服务
func NewService() *Service {
	return &Service{
		db: db.GetInstance(),
	}
}

// GetByID 根据ID获取订单
func (s *Service) GetByID(id uint) (*Order, error) {
	var order Order
	result := s.db.GetDB().Preload("User").First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// GetByUserID 根据用户ID获取订单列表
func (s *Service) GetByUserID(userID uint) ([]Order, error) {
	var orders []Order
	result := s.db.GetDB().Where("user_id = ?", userID).
		Preload("User").
		Order("created_at DESC").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

// Create 创建订单
func (s *Service) Create(order *Order) error {
	return s.db.GetDB().Create(order).Error
}

// Update 更新订单
func (s *Service) Update(order *Order) error {
	return s.db.GetDB().Save(order).Error
}

// UpdateStatus 更新订单状态
func (s *Service) UpdateStatus(id uint, status OrderStatus) error {
	return s.db.GetDB().Model(&Order{}).Where("id = ?", id).Update("status", status).Error
}

// List 获取订单列表
func (s *Service) List(page, pageSize int) ([]Order, int64, error) {
	var orders []Order
	var total int64

	query := s.db.GetDB().Model(&Order{}).Preload("User")
	query.Count(&total)

	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	result := query.Order("created_at DESC").Find(&orders)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return orders, total, nil
}
