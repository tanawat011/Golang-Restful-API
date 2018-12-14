package db

import (
	model "github/Tanawat_DEVz/example-rest-api2/models"
)

var people []model.Person

// Insert allows populating database
func Insert(person model.Person) {
	people = append(people, person)
}

// Get returns the whole database
func Get() []model.Person {
	return people
}
