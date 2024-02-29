package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/CrudOperation/models"
	"github.com/gorilla/mux"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var monuments []models.MonumentAttribute

func main() {
	r := mux.NewRouter()

	monuments = append(monuments, models.MonumentAttribute{Id: "1", Name: "HAWA-MAHAL"})
	monuments = append(monuments, models.MonumentAttribute{Id: "2", Name: "AMER FORT"})

	r.HandleFunc("/monuments", getMonuments).Methods("GET")
	r.HandleFunc("/monument/{id}", getMonument).Methods("GET")
	r.HandleFunc("/monument", createMonument).Methods("POST")
	r.HandleFunc("/monument/{id}", updateMonument).Methods("PUT")
	r.HandleFunc("/monument/{id}", deleteMonument).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func getMonument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range monuments {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func getMonuments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monuments)
}

func createMonument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.MonumentAttribute
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000))
	monuments = append(monuments, book)
	json.NewEncoder(w).Encode(book)
}

func updateMonument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for index, item := range monuments {
		if item.Id == params["id"] {
			monuments = append(monuments[:index], monuments[index+1:]...)
			var book models.MonumentAttribute

			json.NewDecoder(r.Body).Decode(&book)
			book.Id = params["id"]
			monuments = append(monuments, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteMonument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for index, item := range monuments {
		if item.Id == params["id"] {
			monuments = append(monuments[:index], monuments[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(monuments)
}
