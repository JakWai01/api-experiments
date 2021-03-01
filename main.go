package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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
	http.HandleFunc("/", homepage)
	http.HandleFunc("/notes", returnAllNotes)
	log.Fatal(http.ListenAndServe(":10000", nil))
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