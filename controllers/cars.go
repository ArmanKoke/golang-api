package controllers

import (
	"encoding/json"
	"items_api/db"
	"items_api/models"
	"items_api/utils"
	"log"
	"net/http"
)

const idParam = "id"
const statusParam = "status"

func HandleCars(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// only way to check for raw query param
		if rawQueryHasParam(r, idParam) {
			Get(w, r)
		} else {
			GetAll(w, r)
		}
	case "POST":
		Save(w, r)
	case "PUT": //only nominal for now
		Save(w, r)
	case "DELETE":
		Delete(w, r)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w = utils.AddHeaders(w)

	err := json.NewEncoder(w).Encode(db.GetAll())
	if err != nil {
		log.Printf("Error in Get method: %s", err)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	w = utils.AddHeaders(w)
	//not good way to do it but leave for demo
	var body interface{}

	id := utils.ConvertStringId(utils.ProcessRawQuery(r, idParam))

	if rawQueryHasParam(r, statusParam) {
		body = db.GetWithStatus(id)
	} else {
		body = db.Get(id)
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Printf("Error in Get method: %s", err)
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
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

func Delete(w http.ResponseWriter, r *http.Request) {
	w = utils.AddHeaders(w)

	id := utils.ConvertStringId(utils.ProcessRawQuery(r, idParam))

	err := json.NewEncoder(w).Encode(db.Delete(id))
	if err != nil {
		log.Printf("Error in Delete method: %s", err)
	}
}

//not the best way but works for now
func rawQueryHasParam(r *http.Request, param string) bool {
	return len(r.URL.Query()[param]) != 0
}
