package service

import (
	"errors"

	"github.com/msyamsula/messaging-api/user/database"
)

var (
	ErrWrongPassword = errors.New("wrong password")
)

type Service interface {
	Register(username string, password string) (database.User, error)
	Login(username string, password string) (database.User, error) // add token in login
	GetAllUser() ([]database.User, error)
	Logout(username string) error
}
