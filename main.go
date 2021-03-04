package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Accept-Encoding", "Accept-Language", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Cache-Control", "Connection", "Host", "Origin", "Sec-GPC", "User-Agent"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/", homepage)
	r.HandleFunc("/notes", returnAllNotes)
	r.HandleFunc("/note", createNewNote).Methods("POST")
	r.HandleFunc("/article/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/notes/{id}", returnSingleNote)
	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(headers, methods, origins)(r)))
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
		{Id: "1", Title: "Hello World!", Content: "Just another dummy"},
	}
	handleRequests()
}
