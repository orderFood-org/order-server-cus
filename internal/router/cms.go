package router

import (
	"orderFood-server-cus/internal/middleware"
	"orderFood-server-cus/internal/service"

	"github.com/gin-gonic/gin"
)

func CmsRouters(r *gin.Engine, s *service.Service) {
	session := middleware.NewSessionAuth()

	root := r.Group(rootPath).Use(session.Auth)
	{
		root.GET("/cms/hello", s.Hello)
	}
}
