package object

import (
	"github.com/msyamsula/messaging-api/user/database"
	"github.com/msyamsula/messaging-api/user/service"
)

type Service struct {
	Db database.DB
}

func New(db database.DB) (*Service, error) {
	svc := &Service{
		Db: db,
	}

	return svc, nil
}

func (svc *Service) Register(username string, password string) (database.User, error) {
	var err error
	var user database.User

	err = svc.Db.InsertUser(username, password)
	if err != nil {
		return user, err
	}

	user, err = svc.Db.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (svc *Service) Login(username string, password string) (database.User, error) {
	var err error
	var user database.User

	user, err = svc.Db.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	dbPassword := user.Password

	if dbPassword != password {
		return database.User{}, service.ErrWrongPassword
	}

	return user, nil
}

func (svc *Service) GetAllUser() ([]database.User, error) {
	// var err error
	// var users []database.User

	return svc.Db.GetAllUser()
}
