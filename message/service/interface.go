package service

import (
	"context"

	"github.com/msyamsula/messaging-api/message/database"
)

type Service interface {
	InsertMessage(ctx context.Context, m database.MessageToInsert) error
	GetConversation(ctx context.Context, person1 int64, person2 int64) ([]database.Message, error)
	ReadMessage(ctx context.Context, person1 int64, person2 int64) error
}
