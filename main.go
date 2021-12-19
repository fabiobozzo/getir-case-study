package main

import (
	"context"
	"fmt"
	"getir-case-study/api/fetch"
	"getir-case-study/api/inmemory"
	"getir-case-study/pkg"
	"getir-case-study/pkg/db/mongo"
	"net/http"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		logger.Fatalf("error loading .env file: %s", err)
	}

	cfg := pkg.Config{}
	if err := env.Parse(&cfg); err != nil {
		logger.Fatalf("error loading config: %s", err)
	}

	mongoDb, err := mongo.NewMongoDBConnection(cfg.MongoDBURI, cfg.DatabaseName, context.Background())
	if err != nil {
		logger.Fatalf("error connecting to mongodb cluster: %s", err)
	}

	fetchHandler := fetch.NewHandler(mongo.NewReader(mongoDb, cfg.CollectionName))
	inmemoryHandler := inmemory.NewHandler()

	http.HandleFunc("/fetch", fetchHandler.Handle)
	http.HandleFunc("/in-memory", inmemoryHandler.Handle)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatalf("Server terminated ungracefully: %s", err)
	}
}
