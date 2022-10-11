package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	InsertMessage(c *gin.Context)
	ReadMessage(c *gin.Context)
	GetConversation(c *gin.Context)
}
