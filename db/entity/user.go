package entity

type User struct {
	DefaultColumn    `gorm:"embedded"`
	Username         string    `gorm:"unique; not null; column:username"`
	Password         string    `gorm:"not null; column:password;" json:"-"`
	UnreadMessages   int       `gorm:"-"`
	SentMessages     []Message `gorm:"foreignKey:senderID; references:ID" json:"-"`
	ReceivedMessages []Message `gorm:"foreignKey:receiverID; references:ID" json:"-"`
	IsActive         bool      `gorm:"column:isActive; default:0"`
}

func (u User) TableName() string {
	return "users"
}
