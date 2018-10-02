package models

import "github.com/jinzhu/gorm"

type User struct {
	ID uint
	Name string
	Email string
	Email_verified_at string
	Password string
	RememberToken string
	gorm.Model
}
