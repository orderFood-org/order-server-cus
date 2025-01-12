package router

import (
	"orderFood-server-cus/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api"
	noAuthPath = "/out/api"
)

func InitRouters(r *gin.Engine) {
	s := service.NewService()
	CmsRouters(r, s)
	NoAuthRouters(r, s)
}
