package models

import (
	"github.com/jinzhu/gorm"
)

type Migration struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	TableName string `json:"table_name"`
	gorm.Model
}

func InitTables() {

}
