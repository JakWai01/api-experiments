package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A Note consists of a title and some content
type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Notes consists of multiple single Notes
var Notes []Note

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homepage)
	r.HandleFunc("/notes", requests)

	log.Fatal(http.ListenAndServe(":10000", r))
}

func requests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(Notes)
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var note Note
		json.Unmarshal(reqBody, &note)

		Notes = append(Notes, note)
		fmt.Println(w, "%+v", string(reqBody))
	default:
	}
}

func main() {
	handleRequests()
}
