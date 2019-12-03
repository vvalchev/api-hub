package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("ui/"))
	router := NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
	log.Print("Listening on port 10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}
