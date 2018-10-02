package app

import "api/app/service/config"

// load config
// init application
// load middware
// hanlde exception
// init logger handle

func Run() {
	config.Register()
}
