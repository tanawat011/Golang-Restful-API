package model

// Person description.
// swagger:model person
type Person struct {
	// ID of the person
	//
	// required: true
	ID        string   `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

// Address description
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}
