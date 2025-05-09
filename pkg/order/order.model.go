package order

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"orderFood-server-cus/pkg/account"

	"gorm.io/gorm"
)

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"    // 待处理
	OrderStatusProcessing OrderStatus = "processing" // 处理中
	OrderStatusCompleted  OrderStatus = "completed"  // 已完成
	OrderStatusCancelled  OrderStatus = "cancelled"  // 已取消
)

// OrderItem 订单项目
type OrderItem struct {
	DishID uint    `json:"dishId"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Count  int     `json:"count"`
}

// OrderItems 自定义JSON类型用于订单项列表
type OrderItems []OrderItem

// Scan 从数据库中取出JSON数据
func (i *OrderItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("类型断言失败")
	}
	return json.Unmarshal(bytes, i)
}

// Value 存入数据库前将订单项转为JSON
func (i OrderItems) Value() (driver.Value, error) {
	if len(i) == 0 {
		return nil, nil
	}
	return json.Marshal(i)
}

// Order 订单
type Order struct {
	gorm.Model
	UserID         uint        `gorm:"column:user_id;not null" json:"userId"`
	Status         OrderStatus `gorm:"column:status;type:varchar(20);default:'pending';not null" json:"status"`
	TotalAmount    float64     `gorm:"column:total_amount;type:decimal(10,2);not null" json:"totalAmount"`
	Items          OrderItems  `gorm:"column:items;type:json;not null" json:"items"`
	TableNo        string      `gorm:"column:table_no;type:varchar(20);not null" json:"tableNo"`
	NumberOfPeople int         `gorm:"column:number_of_people;not null" json:"numberOfPeople"`
	Remark         string      `gorm:"column:remark;type:varchar(500)" json:"remark"`

	// 关联
	User account.Account `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName 返回表名
func (o *Order) TableName() string {
	return "orders"
}
