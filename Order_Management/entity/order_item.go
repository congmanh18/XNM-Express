package entity

import "time"

type OrderItem struct {
	ID        *string `gorm:"primaryKey"`
	OrderID   *string
	ItemID    *string
	Quantity  *int
	Price     *float64 // suy nghĩ thêm
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o OrderItem) TableName() string {
	return "order_management_service.order_items"
}
