package setting

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

var Config *ini.File
var App = &app{}
var Server = &server{}
var Database = &database{}
var Redis = &redis{}

func Register() {
	var err error

	Config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", App)
	mapTo("server", Server)
	mapTo("database", Database)
	mapTo("redis", Redis)

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.ReadTimeout * time.Second
	Redis.IdleTimeout = Redis.IdleTimeout * time.Second

}

func mapTo(section string, value interface{}) {
	err := Config.Section(section).MapTo(value)

	if err != nil {
		log.Fatalf("Config.MapTo %s Setting err: %v", section, err)
	}
}
