package main

import (
	"log"
	"net/http"

	"github.com/Nikita-Astafyev/api-gateway-go/internal/server"
)

func main() {
	server.StartServer()
	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
