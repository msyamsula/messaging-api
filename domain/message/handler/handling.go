package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/msyamsula/messaging-api/db/entity"
)

// func (h *Handler) CreateMessage(msg entity.Message) (entity.Message, error) {
// 	return h.messageSvc.Create(msg)
// }

func (h *Handler) CreateMessage(c *gin.Context) {

	var msg entity.Message
	var err error
	var status int
	var message string

	c.BindJSON(&msg)

	msg, err = h.messageSvc.Create(msg)

	if err != nil {
		status = http.StatusInternalServerError
		message = FailedInsertMessage
	} else {
		status = http.StatusOK
		message = Success
	}

	response := gin.H{
		"data":    msg,
		"error":   err,
		"message": message,
	}

	c.JSON(status, response)
}

func (h *Handler) GetMessageByUserID(c *gin.Context) {
	var messages []entity.Message
	var err error
	var userID int

	userID, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":  messages,
			"error": err.Error(),
		})
		return
	}

	messages, err = h.messageSvc.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":  messages,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    messages,
		"error":   err,
		"message": Success,
	})
}