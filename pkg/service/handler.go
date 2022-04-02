package service

import (
	"avitotechtask/pkg/database"
	"encoding/json"
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
	var JsonCome JsonCome

	err := json.NewDecoder(r.Body).Decode(&JsonCome)
	if err != nil {
		w.Write([]byte("ERROR : Invalid Json\n"))
	} else {
		if database.QueryProfitMoney(JsonCome.Id, JsonCome.Money) {
			w.Write([]byte("Money added successfully!"))
		}
	}

}

func WriteOffMoneyHandler(w http.ResponseWriter, r *http.Request) {
	var JsonCome1 JsonCome

	err := json.NewDecoder(r.Body).Decode(&JsonCome1)
	if err != nil {
		w.Write([]byte("ERROR : Invalid Json\n"))
	} else {
		if database.QueryWriteOffMoney(JsonCome1.Id, JsonCome1.Money) {
			w.Write([]byte("Money written of successfully!"))
		} else {
			w.Write([]byte("ERROR : Not enough money\n"))
		}
	}

}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	var JsonTransfer JsonTransfer

	err := json.NewDecoder(r.Body).Decode(&JsonTransfer)
	if err != nil {
		w.Write([]byte("ERROR : Invalid Json\n"))
	} else {
		if database.QueryTransfer(JsonTransfer.FirstId, JsonTransfer.SecondId, JsonTransfer.Money) {
			w.Write([]byte("Money transferred!"))
		} else {
			w.Write([]byte("ERROR : Not enough money!\n"))
		}
	}

}

func CheckBalanceHandler(w http.ResponseWriter, r *http.Request) {
	var JsonCome JsonCome

	err := json.NewDecoder(r.Body).Decode(&JsonCome)
	if err != nil {
		w.Write([]byte("ERROR : Invalid Json\n"))
	} else {
		w.Write([]byte(strconv.Itoa(database.GetCurrentBalance(JsonCome.Id))))
	}

}
