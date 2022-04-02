package main

import (
	"avitotechtask/pkg/config"
	"avitotechtask/pkg/database"
	"avitotechtask/pkg/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	service.RouteHandlers()

	database.ConnectDataBase()

	log.Fatalln(http.ListenAndServe(config.GetConfig().Host+config.GetConfig().Port, nil))

}
