package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("starting router")
	router := NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8000", router))
}
