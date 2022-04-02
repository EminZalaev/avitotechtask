package database

import (
	"avitotechtask/pkg/config"
	"database/sql"
	"fmt"
	"log"
)

func ConnectDataBase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.GetConfig().DbHost, config.GetConfig().DbPort, config.GetConfig().User,
		config.GetConfig().Password, config.GetConfig().DbName, config.GetConfig().SSLmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("db dont open", err)
	}

	err = db.Ping()
	if err != nil {
		log.Print(err)
	}

	return db
}
