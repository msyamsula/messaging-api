package service

import "gorm.io/gorm"

type Service struct {
	Db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	svc := &Service{
		Db: db,
	}

	return svc
}
