package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Note struct {
	Id string `json:"Id"`
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
	r.HandleFunc("/notes/{id}", returnSingleNote)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func returnAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllNotes")
	json.NewEncoder(w).Encode(Notes)
}

func returnSingleNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnSingleNote")
	
	vars := mux.Vars(r)
	key := vars["id"]

	for _, note := range Notes {
		if note.Id == key {
			json.NewEncoder(w).Encode(note)
		}
	}
}

func main() {
	Notes = []Note{
		Note{Id: "1", Title: "Hello World!", Content: "Just another dummy"},
	}
	handleRequests()
}