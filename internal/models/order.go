package models

import "gorm.io/gorm"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Total     float64     `json:"total"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
