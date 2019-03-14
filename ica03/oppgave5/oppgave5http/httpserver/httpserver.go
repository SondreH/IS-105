package main

import (
	"fmt"
	"net/http"

	"github.com/karlosmartines/karls-windows/Golang/is105-ica03/oppgave5tcp/jsonemail"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := jsonemail.EncodeJSON()
	fmt.Fprintf(w, string(s))
}

func main() {
	http.HandleFunc("/ok", handler)
	http.ListenAndServe(":5000", nil)
}
