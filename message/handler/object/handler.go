package object

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msyamsula/messaging-api/message/database"
	"github.com/msyamsula/messaging-api/message/handler"
	"github.com/msyamsula/messaging-api/message/service"
)

type Handler struct {
	Svc service.Service
}

func New(s service.Service) handler.Handler {
	h := &Handler{
		Svc: s,
	}

	return h
}

func (h *Handler) InsertMessage(c *gin.Context) {

	var err error
	var ok bool
	var senderID, receiverID int64
	var sid, rid string

	sid, ok = c.GetQuery("senderID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no sender id",
		})
		return
	}

	rid, ok = c.GetQuery("receiverID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no receiver id",
		})
		return
	}

	senderID, err = strconv.ParseInt(sid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	receiverID, err = strconv.ParseInt(rid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	m := database.MessageToInsert{
		SenderID:   senderID,
		ReceiverID: receiverID,
		CreatedAt:  time.Now().Unix(),
		IsRead:     false,
	}

	err = c.BindJSON(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := context.Background()
	err = h.Svc.InsertMessage(ctx, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": m,
	})
}

func (h *Handler) ReadMessage(c *gin.Context) {
	var err error
	var ok bool
	var senderID, receiverID int64
	var sid, rid string

	sid, ok = c.GetQuery("senderID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no sender id",
		})
		return
	}

	rid, ok = c.GetQuery("receiverID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no receiver id",
		})
		return
	}

	senderID, err = strconv.ParseInt(sid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	receiverID, err = strconv.ParseInt(rid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := context.Background()
	err = h.Svc.ReadMessage(ctx, senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})

}
func (h *Handler) GetConversation(c *gin.Context) {
	var err error
	var ok bool
	var senderID, receiverID int64
	var sid, rid string

	sid, ok = c.GetQuery("senderID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no sender id",
		})
		return
	}

	rid, ok = c.GetQuery("receiverID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no receiver id",
		})
		return
	}

	senderID, err = strconv.ParseInt(sid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	receiverID, err = strconv.ParseInt(rid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := context.Background()
	var messages []database.Message
	messages, err = h.Svc.GetConversation(ctx, senderID, receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}
