package object

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	tokenI "github.com/msyamsula/messaging-api/middleware/token"
	"github.com/msyamsula/messaging-api/user/database"
	"github.com/msyamsula/messaging-api/user/handler"
	userSvc "github.com/msyamsula/messaging-api/user/service"
)

type HandlerObj struct {
	Svc userSvc.Service
	t   tokenI.Token
}

func New(svc userSvc.Service, t tokenI.Token) handler.Handler {
	h := &HandlerObj{
		Svc: svc,
		t:   t,
	}

	return h
}

func (h *HandlerObj) Login(c *gin.Context) {
	var err error
	var username string
	var ok bool
	username, ok = c.GetQuery("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no username in http query",
		})
		return
	}

	var password string
	password, ok = c.GetQuery("password")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no password in http query",
		})
		return
	}

	var user database.User
	user, err = h.Svc.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong password",
		})
		return
	}

	var token string
	token, err = h.t.Create(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  user,
		"token": token,
	})
}

func (h *HandlerObj) Register(c *gin.Context) {
	var body handler.LoginBody
	c.BindJSON(&body)

	fmt.Println(body)

	var err error
	var user database.User
	user, err = h.Svc.Register(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *HandlerObj) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *HandlerObj) GetAllUser(c *gin.Context) {
	var err error
	var users []database.User
	users, err = h.Svc.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
