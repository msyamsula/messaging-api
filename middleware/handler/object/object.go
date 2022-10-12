package object

import (
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

	_, err = h.t.Validate(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
