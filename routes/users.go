package routes

import (
	"encoding/json"
	httpModel "github.com/jlee2920/cosmo/structs"
	"net/http"
	"users.git/controller"
)

// CreateUser takes in a handler, extracts all pertinent post data required, and calls the controller to create the user
func (s *Server) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		account := &httpModel.UsersPost{}
		if json.NewDecoder(r.Body).Decode(&account) != nil {
			// Serve JSON Decode error
		}

		newAccount, err := controller.CreateUser(r.Context(), s.DB, account)
		if err != nil {
			// Serve error
		}
		json.NewEncoder(w).Encode(newAccount)
	})
}

// ListUsers takes in a handler, extracts the URL params, and calls the controller to get a list of users
func (s *Server) ListUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filters := &httpModel.ListModel{} // TODO: extract parameters from URL
		listOfAccounts, err := controller.ListUsers(r.Context(), s.DB, filters)
		if err != nil {
			// Serve error
		}
		json.NewEncoder(w).Encode(listOfAccounts)
	})
}