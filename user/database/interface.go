package database

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"username"`
	Password string `json:"-"`
	IsActive bool   `json:"is_active"`
}

type DB interface {
	InsertUser(username string, password string) error
	Login(username string, password string) (User, error)
	GetUserByUsername(username string) (User, error)
	GetAllUser() ([]User, error)
	Logout(username string) error
}
