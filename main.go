package main

import (
	"avitotechtask/pkg/database"
	"avitotechtask/pkg/service/handlers"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	handlers.RouteHandlers()

	database.ConnectDataBase()

	http.ListenAndServe(":8080", nil)

}
