package object

import (
	"context"

	"github.com/msyamsula/messaging-api/message/database"
	"github.com/msyamsula/messaging-api/message/service"
)

type Service struct {
	Db database.Database
}

func New(db database.Database) service.Service {
	svc := &Service{
		Db: db,
	}

	return svc
}

func (svc *Service) InsertMessage(ctx context.Context, m database.MessageToInsert) error {
	return svc.Db.InsertMessage(ctx, m)
}

func (svc *Service) GetConversation(ctx context.Context, person1 int64, person2 int64) ([]database.Message, error) {
	return svc.Db.GetConversation(ctx, person1, person2)
}

func (svc *Service) ReadMessage(ctx context.Context, person1 int64, person2 int64) error {
	return svc.Db.ReadMessage(ctx, person1, person2)
}

func (svc *Service) CountUnread(ctx context.Context, senderID int64, receiverID int64) (int64, error) {
	return svc.Db.CountUnread(ctx, senderID, receiverID)
}
