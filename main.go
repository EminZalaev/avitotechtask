package main

import (
	"avitotechtask/internal/database"
	"avitotechtask/internal/service/handlers"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	handlers.RouteHandlers()

	database.ConnectDataBase()

	http.ListenAndServe(":8080", nil)

}

