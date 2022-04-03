package entity

type User struct {
	DefaultColumn    `gorm:"embedded"`
	Username         string    `gorm:"unique; not null; column:username"`
	Password         string    `gorm:"not null; column:password;"`
	UnreadMessages   int       `gorm:"-"`
	SentMessages     []Message `gorm:"foreignKey:senderID; references:ID"`
	ReceivedMessages []Message `gorm:"foreignKey:receiverID; references:ID"`
}

func (u User) TableName() string {
	return "users"
}
