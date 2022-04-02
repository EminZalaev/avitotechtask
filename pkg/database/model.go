package database

import (
	_ "github.com/lib/pq"
	"log"
)

type Users struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	CardNumber string `json:"CardNumber"`
	Balance    int    `json:"Balance"`
}

func QueryProfitMoney(id int, money int) bool {
	db := ConnectDataBase()

	rows, err := db.Query("UPDATE users SET balance=balance+$2 where id=$1;", id, money)
	if err != nil {
		log.Println("Query Profit error : ", err)
	}

	defer rows.Close()

	return true
}

func QueryWriteOffMoney(id int, money int) bool {
	db := ConnectDataBase()

	if GetCurrentBalance(id) <= money {
		return false
	}

	rows, err := db.Query("UPDATE users SET balance=balance-$2 where id=$1;", id, money)
	if err != nil {
		log.Println("Query Write-Off set balance error:", err)
	}

	defer rows.Close()

	return true
}

func QueryTransfer(id int, id2 int, money int) bool {
	ConnectDataBase()

	if QueryWriteOffMoney(id, money) {
		QueryProfitMoney(id2, money)

		return true
	}

	return false

}

func GetCurrentBalance(id int) int {
	db := ConnectDataBase()

	var uidCurrent Users

	uidCurrent.Id = id

	checkBalance, err := db.Query("SELECT balance FROM users where id=$1", uidCurrent.Id)
	if err != nil {
		log.Println("Query error:", err)
	}

	for checkBalance.Next() {
		if err := checkBalance.Scan(&uidCurrent.Balance); err != nil {
			log.Println("Read from Query Write-Off error:", err)
		}
	}

	return uidCurrent.Balance
}
