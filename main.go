package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Nikita-Astafyev/api-gateway-go/internal/server"
	swagger "github.com/swaggo/http-swagger"
)

func main() {
	http.Handle("/swagger/", swagger.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.StartServer()
	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
