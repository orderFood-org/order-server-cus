package category

import (
	"gorm.io/gorm"
)

// Category 菜品类别
type Category struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(100);not null;uniqueIndex:categories_name_key" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`
	SortOrder   int    `gorm:"column:sort_order;default:0" json:"sortOrder"`
	IsActive    bool   `gorm:"column:is_active;default:true" json:"isActive"`
}

// TableName 返回表名
func (c *Category) TableName() string {
	return "categories"
}
