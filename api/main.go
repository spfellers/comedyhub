package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GOT ALL THESE PEOPLE")
}

// functions
// start with: func
// params can take types e.g. func f(x int)
// specify return type e.g. func f() string { ...
func getPort() string {
	if args := os.Args[1:]; len(args) > 0 {
		return args[0]
	}
	return "8080"
}

// use fmt.Println() for logging
func main() {
	// equals is :=
	port := getPort()
	router := mux.NewRouter().StrictSlash(true)
	// HandleFunc(route, controller)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/people", getPeople)
	log.Fatal(http.ListenAndServe(":" + port, router))
}