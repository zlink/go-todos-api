package config

import (
	"fmt"
	"log"
	"time"
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
var RunMode string
var DebugMode bool

//Register load application config from file
func Register() {
	var err error

	env, err := ini.Load("config/.env.ini")
	if err != nil {
		log.Fatal("Fail to load env.ini")
	}

	fmt.Println(env)

	Env, err := ini.load("config/.env.ini")
	if err != nil {
		log.Fatalf("Fail to load config env.ini file: %v", err)
	}

	Config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	MapTo("app", App)
	MapTo("server", Server)
	MapTo("database", Database)
	MapTo("redis", Redis)

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.ReadTimeout * time.Second
	Redis.IdleTimeout = Redis.IdleTimeout * time.Second

}

//MapTo map config file
func MapTo(section string, value interface{}) {
	err := Config.Section(section).MapTo(value)

	if err != nil {
		log.Fatalf("Config.MapTo %s Setting err: %v", section, err)
	}
}
