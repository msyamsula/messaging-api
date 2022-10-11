package object

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	middlewareI "github.com/msyamsula/messaging-api/middleware/handler"
	tokenI "github.com/msyamsula/messaging-api/middleware/token"
)

type Handler struct {
	t tokenI.Token
}

func New(t tokenI.Token) middlewareI.Handler {
	h := &Handler{
		t: t,
	}

	return h
}

func (h *Handler) ValidateToken(c *gin.Context) {
	var err error

	token := c.GetHeader("x-api-token")

	token, err = h.t.Validate(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	fmt.Println("goes here")
	c.Request.Header.Set("x-api-token", token)

	c.Next()
}
