package token

import (
	"orderFood-server-cus/common/db"
	"time"
)

// Service 令牌服务
type Service struct {
	db *db.Database
}

// NewService 创建令牌服务
func NewService() *Service {
	return &Service{
		db: db.GetInstance(),
	}
}

// GenerateToken 生成新令牌
func (s *Service) GenerateToken(userID uint, tokenType TokenType, expiresIn time.Duration, userAgent, ipAddress string) (*Token, error) {
	// 这里简化了令牌生成逻辑，实际应用中应使用JWT或其他安全机制
	tokenString := generateRandomToken() // 假设有一个生成随机令牌的函数

	token := &Token{
		UserID:    userID,
		Token:     tokenString,
		Type:      tokenType,
		Expires:   time.Now().Add(expiresIn),
		UserAgent: userAgent,
		IPAddress: ipAddress,
	}

	err := s.db.GetDB().Create(token).Error
	if err != nil {
		return nil, err
	}

	return token, nil
}

// VerifyToken 验证令牌
func (s *Service) VerifyToken(tokenString string, tokenType TokenType) (*Token, error) {
	var token Token
	err := s.db.GetDB().Where("token = ? AND type = ? AND blacklisted = ? AND expires > ?",
		tokenString, tokenType, false, time.Now()).
		Preload("User").
		First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// BlacklistToken 将令牌加入黑名单
func (s *Service) BlacklistToken(tokenString string) error {
	return s.db.GetDB().Model(&Token{}).
		Where("token = ?", tokenString).
		Update("blacklisted", true).Error
}

// RemoveExpiredTokens 清理过期令牌
func (s *Service) RemoveExpiredTokens() error {
	return s.db.GetDB().Where("expires < ?", time.Now()).Delete(&Token{}).Error
}

// 生成随机令牌字符串
func generateRandomToken() string {
	// 实际应用中，应该使用更安全的方法生成令牌
	return "token_" + time.Now().Format("20060102150405")
}
