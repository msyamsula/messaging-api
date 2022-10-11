package database

import (
	"context"
	"errors"
)

var (
	ErrCreatingCollection = errors.New("receiver and sender has same id")
)

type Message struct {
	ID         string `bson:"_id" json:"id"`
	SenderID   int64  `bson:"sender_id" json:"sender_id"`
	Text       string `bson:"text" json:"text"`
	ReceiverID int64  `bson:"receiver_id" json:"receiver_id"`
	IsRead     bool   `bson:"is_read" json:"is_read"`
	CreatedAt  int64  `bson:"unix_created_at" json:"unix_created_at"`
}

type MessageToInsert struct {
	SenderID   int64  `bson:"sender_id" json:"sender_id"`
	Text       string `bson:"text" json:"text"`
	ReceiverID int64  `bson:"receiver_id" json:"receiver_id"`
	IsRead     bool   `bson:"is_read" json:"is_read"`
	CreatedAt  int64  `bson:"unix_created_at" json:"unix_created_at"`
}

type Database interface {
	InsertMessage(ctx context.Context, m MessageToInsert) error
	GetConversation(ctx context.Context, person1 int64, person2 int64) ([]Message, error)
	ReadMessage(ctx context.Context, person1 int64, person2 int64) error
}
