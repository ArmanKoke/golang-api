package models

//import "items_api"

type Car struct {
	Id uint32 `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price uint32 `json:"price"`
	Status uint8 `json:"status"`
	Mileage uint32 `json:"mileage"`
}

//Id uint32, Brand string,Model string,Price uint32,Status uint8, Mileage uint32
//func InsertAll() []Car {
//
//	// Insert some cars
//	return []Car{
//		{Id: 1, Brand: "Benz", Model: "c300", Price: 12333, Status:1, Mileage:2000},
//		{Id: 2, Brand: "BMW", Model: "x3", Price: 23444, Status:2, Mileage:3222},
//		{Id: 3, Brand: "Benz", Model: "e500", Price: 12333, Status:1, Mileage:2000},
//		{Id: 4, Brand: "Benz", Model: "c300", Price: 12333, Status:1, Mileage:2000},
//		{Id: 5, Brand: "Benz", Model: "c300", Price: 12333, Status:1, Mileage:2000},
//	}
//}