package main

import (
	"fmt"
	"getir-case-study/api/fetch"
	"getir-case-study/api/inmemory"
	"log"
	"net/http"
)

func main() {
	fetchHandler := fetch.NewHandler()
	inmemoryHandler := inmemory.NewHandler()

	http.HandleFunc("/fetch", fetchHandler.Handle)
	http.HandleFunc("/in-memory", inmemoryHandler.Handle)

	fmt.Printf("Starting REST API server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
