package jsonemail

// Package jsonemail has a struct for a name and an email and a method to fill
// the fields of that structure and encodes it to a JSON stucture

import (
	"encoding/json"
)

// Email is a struct that can contains fields for name and email
type Email struct {
	Name  string
	Email string
}

// EncodeJSON initializes the fields of a Email-struct with authors name and email,
// encodes it to a JSON-structure which it returns to the caller.
func EncodeJSON() []byte {
	email := Email{
		Name:  "Karl Martin Lund",
		Email: "lund.karlmartin@gmail.com",
	}

	var buf []byte
	buf, err := json.Marshal(email)
	if err != nil {
		panic(err)
	}
	return buf
}
