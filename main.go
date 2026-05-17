package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HakanVardarr/go-ecommerce-api/db"
	"github.com/HakanVardarr/go-ecommerce-api/repository"
	"github.com/HakanVardarr/go-ecommerce-api/routes"
)

func main() {

	database, err := db.InitDB("ecommerce.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	productStore := repository.NewProductStore(database)
	userStore := repository.NewUserStore(database)

	// 3. Router'a hem ürün hem kullanıcı deposunu gönderiyoruz
	// routes paketini de güncellemeyi unutma, parametre olarak artık bunları istiyor
	router := routes.NewRouter(productStore, userStore)

	fmt.Println("Server listening on port: 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
