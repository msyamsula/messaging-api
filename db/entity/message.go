package entity

type Message struct {
	DefaultColumn `gorm:"embedded"`
	Text          string `gorm:"column:text;"`
	SenderID      uint   `gorm:"column:senderID"`
	ReceiverID    uint   `gorm:"column:receiverID"`
}

func (u Message) TableName() string {
	return "messages"
}
