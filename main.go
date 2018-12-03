package main

import (
	"api/app/service/config"
	"api/app/service/database"
	"api/routes"
	"fmt"
	"net/http"
	"time"
)

func main() {

	// load config app.ini filefe
	config.Register()

	// register database service
	database.Register()
	defer database.Close()
	//database.DB.AutoMigrate(models.User{})

	// register router service
	router := routes.Register()

	readTimeout := config.Server.ReadTimeout * time.Second
	writeTimeout := config.Server.WriteTimeout * time.Second
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

	// go func() {
	// 	if err := serv.ListenAndServe(); err != nil {
	// 		log.Printf("Listen: %s \n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// Signal.Notify(quit, os.Interrupt)
	// <-quit

	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// defer cancel()

	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown: ", err)
	// }

	// log.Println("Server exiting...")
}
