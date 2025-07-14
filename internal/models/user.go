package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey,uniqueIndex;not null" json:"id"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `json:"email"`
	Password string
}
