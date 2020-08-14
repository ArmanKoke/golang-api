package controllers

import (
	"encoding/json"
	"fmt"
	"items_api/db"
	"items_api/models"
	"items_api/utils"
	"log"
	"net/http"
)

func HandleCars(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAll(w, r)
	case "POST":
		Save(w, r)
	}
}

func HandleCarsWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAll(w, r)
	case "PUT":
		Save(w, r)
	case "DELETE":
		Delete(w, r)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request){
	w = utils.AddHeaders(w)

	err := json.NewEncoder(w).Encode(db.GetAll())
	if err != nil {
		log.Printf("Error in Get method: %s", err)
	}
}

func Get(w http.ResponseWriter, r *http.Request){
	w = utils.AddHeaders(w)

	id := utils.ConvertStringId(utils.ProcessGorillaRequest(r, "id"))

	err := json.NewEncoder(w).Encode(db.Get(id))
	if err != nil {
		log.Printf("Error in Get method: %s", err)
	}
}

func Save(w http.ResponseWriter, r *http.Request){
	w = utils.AddHeaders(w)

	if r.Header.Get("Content-Type") != "" {
		//check content type here to be more strict or better make middleware
	}

	//body content type form-data
	//todo make helper converter
	car := models.Car{
		Id:      utils.ConvertId(r),
		Brand:   r.FormValue("brand"),
		Model:   r.FormValue("model"),
		Price:   utils.ConvertPrice(r),
		Status:  utils.ConvertStatus(r),
		Mileage: utils.ConvertMileage(r),
	}

	err := json.NewEncoder(w).Encode(db.Set(car))
	if err != nil {
		log.Printf("Error in Save method: %s", err)
	}

}

func Delete(w http.ResponseWriter, r *http.Request){
	w = utils.AddHeaders(w)

	id := utils.ConvertStringId(utils.ProcessGorillaRequest(r, "id"))

	a := r.Body
	fmt.Print(a)
	fmt.Print(r.Method)
	err := json.NewEncoder(w).Encode(db.Delete(id))
	if err != nil {
		log.Printf("Error in Delete method: %s", err)
	}
}
