package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	//"github.com/gorilla/mux"
	"items_api/controllers"
	//"items_api/db"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	_,err := fmt.Fprintf(w, "Please check all the routes in README")
	if err != nil {
		log.Printf("Error in homePage method: %s", err)
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cars", controllers.HandleCars)

	//carsRoutes()

	fmt.Println("http://localhost:10001") //for convenience
	log.Fatal(http.ListenAndServe(":10001", nil))
}

func carsRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sadasd")
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Printf("Error in Get method: %s", err)
	}
}

func main() {
	handleRequests()
}
