package service

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	service := &Service{
		db: db,
	}

	return service
}
