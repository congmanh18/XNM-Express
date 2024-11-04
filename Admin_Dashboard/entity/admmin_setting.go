package entity

import "time"

type AdminSetting struct {
	ID        *string `gorm:"primaryKey"`
	Key       *string
	Value     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a AdminSetting) TableName() string {
	return "admin_dashboard_service.admin_settings"
}
