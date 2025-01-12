package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloParams struct {
	Name string `form:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func (s *Service) Hello(c *gin.Context) {
	params := HelloParams{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &HelloResponse{
			Message: fmt.Sprintf("hello %s", params.Name),
		},
	})
}
