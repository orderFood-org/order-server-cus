package token

import (
	"orderFood-server-cus/pkg/account"
	"time"

	"gorm.io/gorm"
)

// TokenType 令牌类型
type TokenType string

const (
	TokenTypeAccess  TokenType = "access"  // 访问令牌
	TokenTypeRefresh TokenType = "refresh" // 刷新令牌
)

// Token 认证令牌
type Token struct {
	gorm.Model
	UserID      uint      `gorm:"column:user_id;not null;index:idx_tokens_user_id" json:"userId"`
	Token       string    `gorm:"column:token;type:varchar(500);not null;uniqueIndex:tokens_token_unique" json:"token"`
	Type        TokenType `gorm:"column:type;type:varchar(20);default:'access';not null;index:idx_tokens_type" json:"type"`
	Expires     time.Time `gorm:"column:expires;not null;index:idx_tokens_expires" json:"expires"`
	Blacklisted bool      `gorm:"column:blacklisted;default:false;not null" json:"blacklisted"`
	UserAgent   string    `gorm:"column:user_agent;type:varchar(500)" json:"userAgent"`
	IPAddress   string    `gorm:"column:ip_address;type:varchar(50)" json:"ipAddress"`

	// 关联
	User account.Account `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName 返回表名
func (t *Token) TableName() string {
	return "tokens"
}
