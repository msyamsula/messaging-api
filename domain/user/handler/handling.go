package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/msyamsula/messaging-api/db/entity"
)

func (h *Handler) UserLogin(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)

	var err error
	user, err = h.userService.Login(user.Username, user.Password)

	var message string
	var status int

	if err != nil {
		status = http.StatusInternalServerError
		message = "login failed"
	} else {
		status = http.StatusOK
		message = "success"
	}

	response := gin.H{
		"data":    user,
		"error":   err,
		"message": message,
	}

	c.JSON(status, response)
}

func (h *Handler) UserRegister(c *gin.Context) {

	newUser := entity.User{}
	c.BindJSON(&newUser)

	var err error
	err = h.userService.Insert(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": "register failed",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    newUser,
		"message": "success",
		"error":   err,
	})
}

func (h *Handler) AllUserGet(c *gin.Context) {
	users, err := h.userService.GetAllUser()
	response := gin.H{
		"data":   users,
		"errors": err,
		"count":  len(users),
	}

	var status int
	if err != nil {
		status = http.StatusInternalServerError
	} else {
		status = http.StatusOK
	}

	c.JSON(status, response)
}

func (h *Handler) GetSpecificUser(c *gin.Context) {
	username := c.Param("id")
	id, cerr := strconv.Atoi(username)

	fmt.Println(username, id)

	var result entity.User
	var err error
	var count int64
	if cerr == nil {
		result, err, count = h.userService.GetUserByID(id)
	} else {
		result, err, count = h.userService.GetUserByUsername(username)
	}

	var message string
	var status int
	if count == int64(0) {
		status = http.StatusNotFound
		message = "user not found"
	} else if err != nil {
		status = http.StatusInternalServerError
		message = "error when getting a specific user"
	} else if err == nil {
		status = http.StatusOK
		message = "success"
	}

	response := gin.H{
		"data":    result,
		"error":   err,
		"message": message,
	}

	c.JSON(status, response)
}
