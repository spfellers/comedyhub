package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GOT ALL THESE PEOPLE")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/people", getPeople)
	log.Fatal(http.ListenAndServe(":8080", router))
}