package handler

import (
	"github.com/gin-gonic/gin"
)

type HandlerName int64

const (
	PongHandler HandlerName = iota
	UserLoginHandler
	UserRegisterHandler
)

type (
	LoginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type Handler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Pong(c *gin.Context)
	GetAllUser(c *gin.Context)
	Logout(c *gin.Context)
}

// func InitUserHandler(svc *service.Service) map[HandlerName]gin.HandlerFunc {
// 	h := map[HandlerName]gin.HandlerFunc{}
// 	ServiceLayer = svc

// 	h[PongHandler] = PongHandlerFunc
// 	h[UserLoginHandler] = UserLoginHandlerFunc

// 	return h
// }
