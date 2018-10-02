package models

type Migration struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Migration string
	batch int
}
