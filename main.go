package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Nikita-Astafyev/api-gateway-go/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.StartServer()
	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
