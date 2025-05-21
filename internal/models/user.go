package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
	Role     string `gorm:"not null;default:USER" json:"role"`
}
