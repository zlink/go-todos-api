package main

import (
	"api/app/service/config"
	"api/database"
	"api/routes"
	"fmt"
	"net/http"
)

func main() {

	// load config app.ini filefe
	config.Register()
	// register database service
	database.Register()
	// register router service
	router := routes.Register()

	readTimeout := config.Server.ReadTimeout
	writeTimeout := config.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	serv := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	serv.ListenAndServe()
}
