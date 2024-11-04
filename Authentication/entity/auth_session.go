package entity

import "time"

type AuthSession struct {
	ID        *string `gorm:"primaryKey"`
	UserID    *string
	Token     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a AuthSession) TableName() string {
	return "authentication_service.auth_sessions"
}
