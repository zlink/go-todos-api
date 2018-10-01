package models

import "github.com/jinzhu/gorm"

type Todos struct {
	ID uint `gorm:"prikary_key;column:"`
	Title string
	Status bool
	gorm.Model
}
