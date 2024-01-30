package server

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next RouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

func AuthenticationMiddleware(next RouteHandler) RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		isAuthenticated := false

		if !isAuthenticated {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
