package main

import (
	"api/app/util/database"
	"api/app/util/setting"
	"api/routes"
	"net/http"
)

func main() {

	setting.Register()
	database.Register()

	router := routes.Register()

	http.ListenAndServe(":8000", router)
}
