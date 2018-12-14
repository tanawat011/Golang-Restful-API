package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	db "github/Tanawat_DEVz/example-rest-api2/db"
	model "github/Tanawat_DEVz/example-rest-api2/models"
)

// IDParam ID
type IDParam struct {
	ID int64 `json:"id"`
}

// Get is an httpHandler for route GET /people
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Get())
}

// Find is an httpHandler for route GET /people/{id}
func Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.Get() {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}

// Create is an httpHandler for route POST /people
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// person := model.Person{}
	// err := json.NewDecoder(r.Body).Decode(&person)
	// if err != nil {
	// 	panic(err)
	// }

	log.Println("r.FormValue => ", r.FormValue(""))
	log.Println("r.FormValue => ", r.Form)
	// log.Println("person => ", person)

	db.Insert(model.Person{ID: "4", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})
	json.NewEncoder(w).Encode(db.Get())
}
