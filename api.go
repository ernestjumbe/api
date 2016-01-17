package main

import (
	"fmt"
	"github.com/gorillla/mux"
	"html"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fm.Fprint(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
