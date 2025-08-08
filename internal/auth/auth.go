package auth

import (
	"net/http"
	"os"
)

// @Summary Логин пользователя
// @Description Генерирует JWT-токен
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string "JWT Token"
// @Router /auth/login [get]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := GenerateToken("test_user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(token))
}

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
