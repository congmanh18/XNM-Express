package entity

import "time"

type Order struct {
	ID         *string `gorm:"primaryKey"`
	UserID     *string
	Status     *string
	Total      *float64
	OrderDate  time.Time
	OrderItems []OrderItem
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o Order) TableName() string {
	return "order_management_service.orders"
}
