package object

import (
	"context"
	"fmt"
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

	type body struct {
		ReceiverID int64  `json:"receiver_id"`
		SenderID   int64  `json:"sender_id"`
		Text       string `json:"text"`
		IsRead     bool   `json:"is_read"`
	}

	b := body{}
	err = c.BindJSON(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error in json" + err.Error(),
		})
		return
	}

	m := database.MessageToInsert{
		SenderID:   b.SenderID,
		ReceiverID: b.ReceiverID,
		Text:       b.Text,
		CreatedAt:  time.Now().Unix(),
		IsRead:     b.IsRead,
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

func (h *Handler) CountUnread(c *gin.Context) {

	var err error

	type params struct {
		SenderID   int64 `form:"senderID"`
		ReceiverID int64 `form:"receiverID"`
	}

	p := params{}
	err = c.Bind(&p)
	fmt.Println(p)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "wrong http params",
		})

		return
	}

	var unread int64
	unread, err = h.Svc.CountUnread(context.Background(), p.SenderID, p.ReceiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": unread,
	})

}
