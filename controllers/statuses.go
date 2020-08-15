package controllers

import (
	"encoding/json"
	"items_api/db"
	"items_api/models"
	"items_api/utils"
	"log"
	"net/http"
)

func HandleStatuses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// only way to check for raw query param
		if rawQueryHasParam(r, idParam) {
			//Get(w, r)
		} else {
			GetAllS(w, r)
		}
	case "POST":
		SaveS(w, r)
	case "PUT":
		SaveS(w, r)
	case "DELETE":
		DeleteS(w, r)
	}
}

func GetAllS(w http.ResponseWriter, r *http.Request) {
	w = utils.AddHeaders(w)

	err := json.NewEncoder(w).Encode(db.GetAllS())
	if err != nil {
		log.Printf("Error in Get method: %s", err)
	}
}

//func Get(w http.ResponseWriter, r *http.Request) {
//	w = utils.AddHeaders(w)
//
//	id := utils.ConvertStringId(utils.ProcessRawQuery(r, idParam))
//
//	err := json.NewEncoder(w).Encode(db.Get(id))
//	if err != nil {
//		log.Printf("Error in Get method: %s", err)
//	}
//}

func SaveS(w http.ResponseWriter, r *http.Request) {
	w = utils.AddHeaders(w)

	if r.Header.Get("Content-Type") != "" {
		//check content type here to be more strict or better make middleware
	}

	//body content type form-data
	//todo make helper converter
	status := models.Status{
		Id:   utils.ConvertId(r),
		Name: r.FormValue("name"),
	}

	err := json.NewEncoder(w).Encode(db.SetS(status))
	if err != nil {
		log.Printf("Error in Save method: %s", err)
	}

}

func DeleteS(w http.ResponseWriter, r *http.Request) {
	w = utils.AddHeaders(w)

	id := utils.ConvertStringId(utils.ProcessRawQuery(r, idParam))

	err := json.NewEncoder(w).Encode(db.DeleteS(id))
	if err != nil {
		log.Printf("Error in Delete method: %s", err)
	}
}
