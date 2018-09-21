package database

import (
	"api/app/util/setting"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Register() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Database.User,
		setting.Database.Password,
		setting.Database.Host,
		setting.Database.Name)
	fmt.Println(dsn)
	db, err = gorm.Open(setting.Database.Type, dsn)
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, table string) string {
		return setting.Database.TablePrefix + table
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func close() {
	defer db.Close()
}
