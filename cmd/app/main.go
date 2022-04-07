package main

import (
	"avitotechtask/pkg/database"
	"avitotechtask/pkg/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	service.RouteHandlers()

	database.ConnectDataBase()
	log.Println("server started")
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
