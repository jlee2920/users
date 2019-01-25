package routes

import (
	"encoding/json"
	httpModel "github.com/jlee2920/cosmo/structs"
	"net/http"
	"users.git/controller"
)

func (s *Server) CreateRole() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := &httpModel.RolePost{}
		if json.NewDecoder(r.Body).Decode(&role) != nil {
			// Serve JSON Decode error
		}

		roleCreate, err := controller.CreateRole(r.Context(), s.DB, role)
		if err != nil {
			// Handle post error
		}
		json.NewEncoder(w).Encode(roleCreate)
	})
}

func (s *Server) ListRoles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filters := &httpModel.ListModel{}
		listOfRoles, err := controller.ListRoles(r.Context(), s.DB, filters)
		if err != nil {
			// Serve error
		}
		json.NewEncoder(w).Encode(listOfRoles)
	})
}
