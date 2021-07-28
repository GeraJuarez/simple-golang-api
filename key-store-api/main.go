package main

import (
	"example/cloud-app/store/registry"
	"example/cloud-app/store/router"
	repository "example/cloud-app/store/usecase/repository/local"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var PORT = "PORT"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	portEnv := os.Getenv(PORT)

	datastore := repository.NewKVStoreLocal()
	registry := registry.NewRegistry(datastore)
	router := router.Start(registry.NewAppController())

	log.Fatal(http.ListenAndServe(":"+portEnv, router))
}
