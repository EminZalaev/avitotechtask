package database

import (
	"avitotechtask/pkg/config"
	"database/sql"
	"fmt"
	"log"
)

type psql struct {
	dbslq *sql.DB
}

func ConnectDataBase() *sql.DB {
	var db *sql.DB
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.GetConfig().User,
		config.GetConfig().Password,
		config.GetConfig().DbHost,
		config.GetConfig().DbPort,
		config.GetConfig().DbName)

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
