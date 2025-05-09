package dish

import (
	"orderFood-server-cus/pkg/category"

	"gorm.io/gorm"
)

// Dish 菜品
type Dish struct {
	gorm.Model
	Name        string  `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Description string  `gorm:"column:description;type:text" json:"description"`
	Price       float64 `gorm:"column:price;type:decimal(10,2);not null" json:"price"`
	SoldCount   int     `gorm:"column:sold_count;default:0" json:"soldCount"`
	ImageURL    string  `gorm:"column:image_url;type:varchar(255)" json:"imageUrl"`
	CategoryID  uint    `gorm:"column:category_id;not null" json:"categoryId"`
	IsSpecial   bool    `gorm:"column:is_special;default:false" json:"isSpecial"`
	IsAvailable bool    `gorm:"column:is_available;default:true" json:"isAvailable"`
	SortOrder   int     `gorm:"column:sort_order;default:0" json:"sortOrder"`

	// 关联
	Category category.Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// TableName 返回表名
func (d *Dish) TableName() string {
	return "dishes"
}
