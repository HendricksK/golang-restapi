package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
}

//main func
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Ping).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3030", router))
}

var people []Person

func Ping(w http.ResponseWriter, r *http.Request) {
	// The person Type (more like an object)
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe"})
	json.NewEncoder(w).Encode(people)
}
func GetPeople(w http.ResponseWriter, r *http.Request) {}
func GetPerson(w http.ResponseWriter, r *http.Request) {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}