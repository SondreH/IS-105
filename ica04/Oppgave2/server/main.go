package main

import (


	"net/http"
	"html/template"
	"path"

)

func main() {


	http.HandleFunc("/", foo)
	http.HandleFunc("/speech", redirect)
	http.ListenAndServe(":80", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {


	side := path.Join("html","index.html")
	html, err := template.ParseFiles(side)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := html.Execute(w, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func redirect(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	text := values.Get("text")
	sprok := values.Get("sprok")
	pitch := values.Get("pitch")
	speed := values.Get("speed")


	if sprok == "Norsk" {
		http.Redirect(w, r, "http://0.0.0.0:8080/speech?text="+text+"&voice=no&pitch="+pitch+"&speed="+speed, 301)
		} else if sprok == "Engelsk" {
		http.Redirect(w, r, "http://0.0.0.0:8080/speech?text="+text+"&voice=en&pitch="+pitch+"&speed="+speed, 301)
		} else if sprok == "Russisk" {
		http.Redirect(w, r, "http://0.0.0.0:8080/speech?text="+text+"&voice=sr&pitch="+pitch+"&speed="+speed, 301)
	}
}
