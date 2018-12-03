package main

import (
	"api/app/service/config"
	"api/database"
	"api/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Printf("Listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal)
	Signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting...")
}
