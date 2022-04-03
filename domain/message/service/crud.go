package service

import (
	"github.com/msyamsula/messaging-api/db/entity"
)

func (s *Service) Create(msg entity.Message) (entity.Message, error) {

	trx := s.Db.Begin()

	t := trx.Create(&msg)
	if t.Error != nil {
		return entity.Message{}, ErrOperations
	}

	t = trx.Commit()
	if t.Error != nil {
		t.Rollback()
		return entity.Message{}, ErrCommitDB
	}

	return msg, nil
}

func (s *Service) Get(senderID int, receiverID int) ([]entity.Message, error) {
	var msgs []entity.Message

	t := s.Db.Where(&entity.Message{
		SenderID:   uint(senderID),
		ReceiverID: uint(receiverID),
	}).Or(&entity.Message{
		SenderID:   uint(receiverID),
		ReceiverID: uint(senderID),
	}).Order("created_at asc").Find(&msgs)

	if t.Error != nil {
		return []entity.Message{}, t.Error
	}

	return msgs, nil
}

func (s *Service) ReadMessages(senderID int, receiverID int) error {

	var unreadMessages []entity.Message
	var err error
	var receiver entity.User

	trx := s.Db.Begin()

	t := trx.Where(&entity.User{
		DefaultColumn: entity.DefaultColumn{
			ID: uint(receiverID),
		},
	}).Find(&receiver)
	if t.Error != nil {
		trx.Rollback()
		return t.Error
	}

	err = trx.Model(&receiver).Where(map[string]interface{}{
		"isRead":   false,
		"senderID": senderID,
	}).Association("ReceivedMessages").Find(&unreadMessages)
	if err != nil {
		trx.Rollback()
		return err
	}

	for _, um := range unreadMessages {
		um.IsRead = true
		t = trx.Save(&um)
		if t.Error != nil {
			trx.Rollback()
			return t.Error
		}
	}

	trx.Commit()
	return nil
}
