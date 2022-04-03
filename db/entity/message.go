package entity

type Message struct {
	DefaultColumn `gorm:"embedded"`
	Text          string `gorm:"column:text;"`
	SenderID      uint   `gorm:"column:senderID"`
	ReceiverID    uint   `gorm:"column:receiverID"`
	IsRead        bool   `gorm:"column:isRead; default:0"`
}

func (u Message) TableName() string {
	return "messages"
}
