package account

import "gorm.io/gorm"

// UserRole 用户角色
type UserRole string

const (
	UserRoleAdmin UserRole = "admin" // 管理员
	UserRoleStaff UserRole = "staff" // 员工
	UserRoleUser  UserRole = "user"  // 普通用户
)

// Account 用户账户
type Account struct {
	gorm.Model
	Username string   `gorm:"column:username;type:varchar(50);not null;uniqueIndex:users_username_unique" json:"username"`
	Email    string   `gorm:"column:email;type:varchar(100);not null;uniqueIndex:users_email_unique" json:"email"`
	Password string   `gorm:"column:password;type:varchar(255);not null" json:"password,omitempty"`
	FullName string   `gorm:"column:full_name;type:varchar(100)" json:"fullName"`
	Role     UserRole `gorm:"column:role;type:varchar(20);default:'user';not null" json:"role"`
	IsActive bool     `gorm:"column:is_active;default:true;not null" json:"isActive"`
}

// TableName 返回表名
func (a *Account) TableName() string {
	return "users"
}
