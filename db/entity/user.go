package entity

type User struct {
	DefaultColumn `gorm:"embedded"`
	Username      string `gorm:"unique; not null; column:username"`
	Password      string `gorm:"not null; column:password;"`
}

func (u User) TableName() string {
	return "users"
}
