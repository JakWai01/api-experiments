package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

var Notes []Note

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("homepage")
}

func handleRequests() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homepage)
	r.HandleFunc("/notes", returnAllNotes)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func returnAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllNotes")
	json.NewEncoder(w).Encode(Notes)
}

func main() {
	Notes = []Note{
		Note{Title: "Hello World!", Content: "Just another dummy"},
	}
	handleRequests()
}