package auth

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func APIKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		validKey := os.Getenv("API_KEY")

		if apiKey != validKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	}
}
