package main

import (
	"fmt"
	"getir-case-study/api/fetch"
	"getir-case-study/api/inmemory"
	"getir-case-study/pkg"
	"log"
	"net/http"

	"github.com/caarlos0/env"
)

func main() {
	cfg := pkg.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to load config")
	}

	fetchHandler := fetch.NewHandler()
	inmemoryHandler := inmemory.NewHandler()

	http.HandleFunc("/fetch", fetchHandler.Handle)
	http.HandleFunc("/in-memory", inmemoryHandler.Handle)

	fmt.Printf("Starting REST API server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
