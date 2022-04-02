package handlers

import (
	"avitotechtask/pkg/database"
	"avitotechtask/pkg/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func RouteHandlers() {
	http.HandleFunc("/BillingService/ProfitMoney", ProfitMoneyHandler)
	http.HandleFunc("/BillingService/WriteOffMoney", WriteOffMoneyHandler)
	http.HandleFunc("/BillingService/Transfer", TransferHandler)
	http.HandleFunc("/BillingService/CheckBalance", CheckBalanceHandler)
}

func ProfitMoneyHandler(w http.ResponseWriter, r *http.Request) {
	var JsonCome service.JsonCome

	err := json.NewDecoder(r.Body).Decode(&JsonCome)
	if err != nil {
		log.Println("ERR :cant decode", err)
	}

	database.QueryProfitMoney(JsonCome.Id, JsonCome.Money)

	w.Write([]byte("Money added successfully!"))
}

func WriteOffMoneyHandler(w http.ResponseWriter, r *http.Request) {
	var JsonCome1 service.JsonCome

	err := json.NewDecoder(r.Body).Decode(&JsonCome1)
	if err != nil {
		log.Println("ERR :cant decode", err)
	}

	if !database.QueryWriteOffMoney(JsonCome1.Id, JsonCome1.Money) {
		w.Write([]byte("ERROR : Not enough money"))
	}

	w.Write([]byte("Money written of successfully!"))
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	var JsonTransfer service.JsonTransfer

	err := json.NewDecoder(r.Body).Decode(&JsonTransfer)
	if err != nil {
		log.Println("ERR :cant decode", err)
	}

	database.QueryTransfer(JsonTransfer.FirstId, JsonTransfer.SecondId, JsonTransfer.Money)

	w.Write([]byte("Money transferred!"))
}

func CheckBalanceHandler(w http.ResponseWriter, r *http.Request) {
	var JsonCome service.JsonCome

	err := json.NewDecoder(r.Body).Decode(&JsonCome)
	if err != nil {
		log.Println("ERR :cant decode", err)
	}

	w.Write([]byte(strconv.Itoa(database.GetCurrentBalance(JsonCome.Id))))
}
