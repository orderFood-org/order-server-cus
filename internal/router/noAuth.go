package router

import (
	"orderFood-server-cus/internal/service"

	"github.com/gin-gonic/gin"
)

func NoAuthRouters(r *gin.Engine, s *service.Service) {
	noAuth := r.Group(noAuthPath)
	// /out/api
	{
		noAuth.POST("/cms/register", s.Register)
		noAuth.POST("/cms/login", s.Login)
	}
}
