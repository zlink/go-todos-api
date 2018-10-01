package database

import (
	"api/app/service/config"
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB  *gorm.DB
	err error
)

func Register() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	config.Database.User,
	//	config.Database.Password,
	//	config.Database.Host,
	//	config.Database.Name,
	//)

	// DB, err = gorm.Open(config.Database.Type, dsn)
	DB, err = gorm.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, table string) string {
		return config.Database.TablePrefix + table
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func close() {
	defer DB.Close()
}
