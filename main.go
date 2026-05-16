package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HakanVardarr/go-ecommerce-api/repository"
	"github.com/HakanVardarr/go-ecommerce-api/routes"
)

func main() {
	store, err := repository.NewProductStore("products.json")
	if err != nil {
		log.Fatalf("Failed to start store: %v", err)
	}

	router := routes.NewRouter(store)

	fmt.Println("Server listening on port: 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
