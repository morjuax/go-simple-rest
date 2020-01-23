package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people[]Person


func getPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func storePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var person Person 
	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = params["id"]
	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}

func updatePersonEndPoint(w http.ResponseWriter, req *http.Request) {


}

func deletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "juan", LastName: "Moreno", Address: 
	&Address{City: "Santiago", State: "Santiago"}})
	
	people = append(people, Person{ID: "2", FirstName: "Pedro", LastName: "Perez", Address: 
	&Address{City: "Santiago", State: "Santiago"}})

	// endpoints
	router.HandleFunc("/people", getPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", storePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/", updatePersonEndPoint).Methods("PUT")
	router.HandleFunc("/people/{id}", updatePersonEndPoint).Methods("DELETE")

	http.ListenAndServe(":3000", router)

	log.Fatal(http.ListenAndServe(":3000", router))

}