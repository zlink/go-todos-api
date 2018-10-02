package models

import "github.com/jinzhu/gorm"

type PasswordResets struct {
	Email string
	Token string
	gorm.Model
}
