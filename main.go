package main

import (
	"context"
	"fmt"
	"getir-case-study/api/fetch"
	"getir-case-study/api/inmemory"
	"getir-case-study/pkg"
	"getir-case-study/pkg/db/mongo"
	"getir-case-study/pkg/kv"
	"net/http"
	"os"

	_ "getir-case-study/docs"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title Getir Coding Assignment
// @version 1.0
// @description The case study for the position of Go developer @Getir
// @contact.name Fabio Bozzo
// @contact.email fabio.bozzo@gmail.com
func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		logger.Warnf("error loading .env file: %s", err)
	}

	cfg := pkg.Config{}
	if err := env.Parse(&cfg); err != nil {
		logger.Fatalf("error loading config: %s", err)
	}

	mongoDb, err := mongo.NewMongoDBConnection(cfg.MongoDBURI, cfg.DatabaseName, context.Background())
	if err != nil {
		logger.Fatalf("error connecting to mongodb cluster: %s", err)
	}

	kvStorage := kv.NewMapStorage()

	fetchHandler := fetch.NewHandler(logger, mongo.NewReader(mongoDb, cfg.CollectionName))
	inmemoryHandler := inmemory.NewHandler(logger, kvStorage)

	http.HandleFunc("/fetch", fetchHandler.Handle)
	http.HandleFunc("/in-memory", inmemoryHandler.Handle)

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		logger.Fatalf("Server terminated ungracefully: %s", err)
	}
}
