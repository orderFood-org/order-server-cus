package service

import (
	"context"
	"fmt"
	"net/http"
	"orderFood-server-cus/internal/dao"
	"orderFood-server-cus/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
	NickName  string `json:"nick_name"`
}

func (s *Service) Login(c *gin.Context) {
	params := LoginParams{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountDao := dao.NewAccountDao(s.db)
	account, err := accountDao.GetAccountByUsername(params.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号不存在"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(params.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确"})
		return
	}

	sessionID, err := s.generateSessionID(context.Background(), params.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LoginResponse{
			SessionID: sessionID,
			Username:  account.Username,
			NickName:  account.NickName,
		},
	})
}

func (s *Service) generateSessionID(ctx context.Context, username string) (string, error) {
	sessionID := uuid.New().String()
	sessionKey := utils.GetSessionKey(username)

	err := s.rdb.Set(ctx, sessionKey, sessionID, time.Hour*8).Err()
	if err != nil {
		fmt.Printf("rdb set session Error: %v", err)
		return "", err
	}

	authKey := utils.GetAuthKey(sessionID)
	err = s.rdb.Set(ctx, authKey, time.Now().Unix(), time.Minute).Err()
	if err != nil {
		fmt.Printf("rdb set session auth Error: %v", err)
		return "", err
	}

	return sessionID, nil
}
