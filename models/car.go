package models

//import "items_api"

type Car struct {
	Id      uint32 `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  uint8  `json:"status"`
	Mileage uint32 `json:"mileage"`
}

type CarWithStatus struct {
	Id      uint32 `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint32 `json:"price"`
	Status  Status `json:"status"`
	Mileage uint32 `json:"mileage"`
}

func DummyCars() []Car {
	return []Car{
		{Id: 1, Brand: "Benz", Model: "c300", Price: 12333, Status: 1, Mileage: 2000},
		{Id: 2, Brand: "BMW", Model: "x3", Price: 23444, Status: 2, Mileage: 3222},
		{Id: 3, Brand: "Benz", Model: "e500", Price: 12333, Status: 3, Mileage: 2000},
		{Id: 4, Brand: "Benz", Model: "c300", Price: 12333, Status: 4, Mileage: 2000},
		{Id: 5, Brand: "Benz", Model: "c300", Price: 12333, Status: 1, Mileage: 2000},
	}
}

func DummyCarsWithStatuses() []CarWithStatus {
	return []CarWithStatus{
		{Id: 1, Brand: "Benz", Model: "c300", Price: 12333, Status: Status{Id: 1, Name: "В пути"}, Mileage: 2000},
		{Id: 2, Brand: "BMW", Model: "x3", Price: 23444, Status: Status{Id: 2, Name: "На складе"}, Mileage: 3222},
		{Id: 3, Brand: "Benz", Model: "e500", Price: 12333, Status: Status{Id: 3, Name: "Продан"}, Mileage: 2000},
		{Id: 4, Brand: "Benz", Model: "c300", Price: 12333, Status: Status{Id: 4, Name: "Снят с продажи"}, Mileage: 2000},
		{Id: 5, Brand: "Benz", Model: "c300", Price: 12333, Status: Status{Id: 1, Name: "В пути"}, Mileage: 2000},
	}
}
