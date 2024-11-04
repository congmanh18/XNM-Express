package entity

import "time"

type Delivery struct {
	ID           *string `gorm:"primaryKey"`
	OrderID      *string
	Status       *string
	DeliveryDate time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (d Delivery) TableName() string {
	return "delivery_tracking_service.deliveries"
}
