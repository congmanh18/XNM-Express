package entity

import "time"

type User struct {
	ID           *string `gorm:"primaryKey"`
	Username     *string
	PasswordHash *string
	Email        *string
	Phone        *string
	Role         *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u User) TableName() string {
	return "user_profile_service.users"
}
