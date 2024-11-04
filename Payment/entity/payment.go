package entity

import "time"

type Payment struct {
	ID          *string `gorm:"primaryKey"`
	OrderID     *string
	Amount      *float64
	Status      *string
	PaymentDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p Payment) TableName() string {
	return "payment_service.payments"
}
