package entity

import "time"

type Inventory struct {
	ID          *string `gorm:"primaryKey"`
	Name        *string
	Description *string
	Quantity    *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i Inventory) TableName() string {
	return "inventory_service.inventories"
}
