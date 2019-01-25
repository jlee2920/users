package middleware

import (
	"io"
	"net/http"
	"strings"
)

var MyAPIKey = "12345"

// Have a general Middleware func to handle a list of middleware for a given route -- need to generalize and make a library so multiple
// services can use this middleware (DRY)
func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

// Example Authentication Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestKey := r.Header.Get("key")
		if len(requestKey) == 0 || strings.Compare(requestKey, MyAPIKey) != 0 {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"invalid_key"}`)
			return
		}

		next.ServeHTTP(w, r)
	})
}