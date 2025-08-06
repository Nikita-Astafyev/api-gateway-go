package server

import (
	"net/http"

	"github.com/Nikita-Astafyev/api-gateway-go/internal/auth"
)

func StartServer() {
	http.HandleFunc("/api/users", userHandler)
	http.HandleFunc("/api/products", productHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	auth.APIKeyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Service (TODO)"))
	})(w, r)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Product Service (TODO: Forward to real service)"))
}
