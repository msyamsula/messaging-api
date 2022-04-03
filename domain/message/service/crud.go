package service

import "github.com/msyamsula/messaging-api/db/entity"

func (s *Service) Create(msg entity.Message) (entity.Message, error) {

	trx := s.Db.Begin()

	t := trx.Create(&msg)
	if t.Error != nil {
		return entity.Message{}, ErrOperations
	}

	t = trx.Commit()
	if t.Error != nil {
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
