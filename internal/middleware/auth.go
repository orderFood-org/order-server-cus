package middleware

import (
	"context"
	"net/http"
	"orderFood-server-cus/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const SessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{}
	connectRDB(s)
	return s
}

// session鉴权
func (s *SessionAuth) Auth(c *gin.Context) {
	sessionID := c.GetHeader(SessionKey)
	if sessionID == "" {
		// AbortWithStatusJSON只是不执行后续的中间件,但是会把本函数执行完,可能会出现多个返回文本,所以直接加return
		c.AbortWithStatusJSON(http.StatusForbidden, "session is null")
		return
	}

	authKey := utils.GetAuthKey(sessionID)
	loginTime, err := s.rdb.Get(c, authKey).Result()
	if err != nil && err != redis.Nil { // 排除空值报错, 真的出现错误
		c.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
		return
	}
	if loginTime == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "session auth fail")
		return
	}

	c.Next()
}

func connectRDB(s *SessionAuth) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	s.rdb = rdb
}
