package main

import (
	"fmt"
	"items_api/controllers"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Please check all the routes in README")
	if err != nil {
		log.Printf("Error in homePage method: %s", err)
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/cars", controllers.HandleCars)
	http.HandleFunc("/statuses", controllers.HandleStatuses)

	fmt.Println("http://localhost:10001") //for convenience
	log.Fatal(http.ListenAndServe(":10001", nil))
}

func main() {
	handleRequests()
}
