package auth

import (
	"net/http"
	"os"
)

func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != os.Getenv("API_KEY") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid API key"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
