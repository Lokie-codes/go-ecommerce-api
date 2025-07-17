package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
