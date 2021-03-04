package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Note struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
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
	r.HandleFunc("/note", createNewNote).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/article/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/notes/{id}", returnSingleNote)
	r.Use(mux.CORSMethodMiddleware(r))
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

func createNewNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("note"))

	fmt.Println("createNewNote")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var note Note
	json.Unmarshal(reqBody, &note)

	Notes = append(Notes, note)
	fmt.Println(w, "%+v", string(reqBody))
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, note := range Notes {

		if note.Id == id {
			Notes = append(Notes[:index], Notes[index+1:]...)
		}
	}
}

func main() {
	Notes = []Note{
		Note{Id: "1", Title: "Hello World!", Content: "Just another dummy"},
	}
	handleRequests()
}
