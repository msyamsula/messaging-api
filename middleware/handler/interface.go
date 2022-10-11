package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrNoToken = errors.New("no token in request")
)

type Handler interface {
	ValidateToken(c *gin.Context)
}
