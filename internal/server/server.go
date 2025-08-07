package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func userProxy() http.Handler {
	target, _ := url.Parse("http://localhost:8081")
	return httputil.NewSingleHostReverseProxy(target)
}

func StartServer() {
	http.Handle("/api/users", userProxy())
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
