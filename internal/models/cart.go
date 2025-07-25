package models

import "gorm.io/gorm"

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
