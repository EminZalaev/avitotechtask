package service

type JsonCome struct {
	Id    int `json:"Id"`
	Money int `json:"Money"`
}

type JsonTransfer struct {
	FirstId  int `json:"FirstId"`
	SecondId int `json:"SecondId"`
	Money    int `json:"Money"`
}

type JsonBalance struct {
	Money int `json:"Money"`
}
