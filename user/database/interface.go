package database

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

type User struct {
	ID       int64
	Name     string
	Password string
	IsActive string
}

type DB interface {
	InsertUser(username string, password string) error
	GetUserByID(id string) (User, error)
	GetUserByUsername(username string) (User, error)
	GetAllUser() ([]User, error)
}
