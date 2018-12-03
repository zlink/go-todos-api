package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type app struct {
	RunMode     string
	PageSize    int
	Secret      string
	RuntimePath string
	TimeFormat  string
}

type server struct {
	HttpPort     int
	SocketPort   int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type redis struct {
	Host        string
	Password    int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type env struct {
	Env string
}

var Config *ini.File
var App = &app{}
var Server = &server{}
var Database = &database{}
var Redis = &redis{}

//Register load application config from file
func Register() {

	var err error

	// Config, err = ini.Load("config/.env")
	// if err != nil {
	// 	log.Fatalf("load env file %t, %v", err, err)
	// }

	Config, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/config.ini': %v", err)
	}

	Config.Section("app").MapTo(App)
	Config.Section("server").MapTo(Server)
	Config.Section("database-ymaccount").MapTo(Database)
}
