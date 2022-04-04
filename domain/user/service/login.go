package service

import (
	"github.com/msyamsula/messaging-api/db/entity"
)

func (s *Service) Login(username string, password string) (entity.User, error) {
	var user entity.User

	db := s.db.Where("username = ? AND password = ?", username, password).First(&user)

	trx := s.db.Begin()
	user.IsActive = true
	t := trx.Model(&user).Update("isActive", true)

	if t.Error != nil {
		return user, t.Error
	}
	t = trx.Commit()
	if t.Error != nil {
		return user, t.Error
	}

	return user, db.Error
}

func (s *Service) Logout(userID int) error {
	var user entity.User

	db := s.db.Where("id = ?", userID).First(&user)

	trx := s.db.Begin()
	user.IsActive = true
	t := trx.Model(&user).Update("isActive", false)

	if t.Error != nil {
		return t.Error
	}
	t = trx.Commit()
	if t.Error != nil {
		return t.Error
	}

	return db.Error
}
