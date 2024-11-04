package entity

import "time"

type Report struct {
	ID        *string `gorm:"primaryKey"`
	Type      *string
	Data      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r Report) TableName() string {
	return "reporting_service.reports"
}
