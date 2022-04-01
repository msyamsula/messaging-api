package entity

import "time"

type DefaultColumn struct {
	ID        uint      `gorm:"column:id; primaryKey; autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at; not null; autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at; not null; autoCreateTime; autoUpdatetime:mili"`
}
