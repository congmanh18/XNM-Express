package entity

import "time"

type Notification struct {
	ID        *string `gorm:"primaryKey"`
	UserID    *string
	Message   *string
	Status    *string
	SentDate  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n Notification) TableName() string {
	return "notification_service.notifications"
}
