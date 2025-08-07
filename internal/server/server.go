package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Nikita-Astafyev/api-gateway-go/internal/auth"
)

func userProxy() http.Handler {
	target, _ := url.Parse("http://localhost:8081")
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Director = func(r *http.Request) {
		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host
		r.URL.Path = "/"
		r.Host = target.Host
	}

	return proxy
}

func productProxy() http.Handler {
	target, _ := url.Parse("http://localhost:8082")
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func StartServer() {
	http.Handle("/api/users", auth.APIKeyMiddleware(userProxy()))
	http.Handle("/api/products", auth.APIKeyMiddleware(productProxy()))
}

/*func userHandler(w http.ResponseWriter, r *http.Request) {
	auth.APIKeyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User Service (TODO)"))
	})(w, r)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Product Service (TODO: Forward to real service)"))
}*/
