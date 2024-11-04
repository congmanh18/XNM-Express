package entity

import "time"

type AuditLog struct {
	ID        *string `gorm:"primaryKey"`
	UserID    *string
	Action    *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a AuditLog) TableName() string {
	return "audit_log_service.audit_logs"
}
