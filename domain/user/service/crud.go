package service

import (
	"sort"

	"github.com/msyamsula/messaging-api/db/entity"
	"gorm.io/gorm"
)

func (s *Service) Insert(user *entity.User) error {
	trx := s.db.Begin()

	var db *gorm.DB

	db = trx.Create(user)
	if db.Error != nil {
		return db.Error
	}

	db = trx.Commit()
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (s *Service) GetAllUser(activeID int) ([]entity.User, error) {
	var users []entity.User
	db := s.db.Find(&users)

	var result []entity.User
	for _, u := range users {
		unread := s.db.Model(&u).Where(map[string]interface{}{
			"receiverID": uint(activeID),
			"isRead":     false,
		}).Association("SentMessages").Count()

		u.UnreadMessages = int(unread)
		result = append(result, u)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].UnreadMessages > result[j].UnreadMessages
	})

	return result, db.Error
}

func (s *Service) GetUserByID(id int) (entity.User, error, int64) {
	var result entity.User
	db := s.db.Where("id = ?", id).Find(&result)

	return result, db.Error, db.RowsAffected
}

func (s *Service) GetUserByUsername(username string) (entity.User, error, int64) {
	var result entity.User
	db := s.db.Where("username = ?", username).Find(&result)

	return result, db.Error, db.RowsAffected
}
