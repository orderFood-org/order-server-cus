package service

import (
	"fmt"
	"net/http"
	"orderFood-server-cus/internal/dao"
	"orderFood-server-cus/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func (s *Service) Register(c *gin.Context) {
	params := RegisterParams{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 密码加密
	hashedPassword, err := encryptPassword(params.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("hashedPassword: %v", hashedPassword)
	// 账号是否存在
	accountDao := dao.NewAccountDao(s.db)
	isExist, err := accountDao.IsExist(params.Username)
	if isExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account is exist"})
		return
	}
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 账号入库
	err = accountDao.Create(&model.Account{
		Username: params.Username,
		Password: hashedPassword,
		NickName: params.NickName,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &RegisterResponse{
			Message: "register success",
		},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password Error: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}
