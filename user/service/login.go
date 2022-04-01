package service

import "github.com/msyamsula/messaging-api/db/entity"

func (s *Service) Login(username string, password string) (entity.User, error) {
	var user entity.User

	db := s.db.Where("username = ? AND password = ?", username, password).First(&user)

	return user, db.Error
}
