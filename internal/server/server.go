package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Nikita-Astafyev/api-gateway-go/internal/auth"
)

func userProxy() http.Handler {
	target, _ := url.Parse("http://localhost:8081")
	return httputil.NewSingleHostReverseProxy(target)
}

func StartServer() {
	http.HandleFunc("/auth/login", auth.LoginHandler)
	http.Handle("/api/users", userProxy())
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
