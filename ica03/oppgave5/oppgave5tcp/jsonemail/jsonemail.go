package jsonemail

import (
	"encoding/json"
)

// Email is a struct that contains name and email
type Email struct {
	Name  string
	Email string
}

// EncodeJSON encodes JSON
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
