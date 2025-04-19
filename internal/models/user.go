package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
	RoleID   uint   `gorm:"not null"'`
	Role     Role   `gorm:"foreignkey:RoleID" json:"role"`
}
