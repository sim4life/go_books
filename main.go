package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	http.HandleFunc("/", PresentBooks)
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	js, err := json.MarshalIndent(book, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func PresentBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building scalable apps with Go", "Shah Ismail"}

	enc := json.NewEncoder(w)
	if err := enc.Encode(&book); err != nil {
		fmt.Println(err)
	}
}
