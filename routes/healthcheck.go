package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"users.git/controller"
)

// CreateUser takes in a handler, extracts all pertinent post data required, and calls the controller to create the user
func (s *Server) Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newAccount, err := controller.Ping(r.Context())
		if err != nil {
			// Serve error
		}
		json.NewEncoder(w).Encode(newAccount)
	})
}

func (s *Server) Accept() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		fmt.Printf("%s\n", buf.String())
		json.NewEncoder(w).Encode(buf.String())
	})
}