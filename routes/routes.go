package routes

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"users.git/config"
)

type Server struct {
	Router *mux.Router
	Config config.Conf
	DB *gorm.DB
}

// Note: Using Handle instead of HandleFunc allows us to use Middleware on a specific route instead of on all of them
// in main.go
func (s *Server) InitializeRoutes() {
	r := mux.NewRouter().StrictSlash(true)

	// Healthcheck routes
	r.Handle("/api/v2/ping", s.Ping()).Methods("POST")
	r.Handle("/api/v2/accept", s.Accept()).Methods("POST")

	// Users routes
	r.Handle("/api/v2/users", s.CreateUser()).Methods("POST")
	r.Handle("/api/v2/users", s.ListUsers()).Methods("GET")
	//r.HandleFunc("/api/v2/accounts/{id}", s.GetAccount).Methods("GET")
	//r.HandleFunc("/api/v2/accounts/{id}", s.ModifyAccount).Methods("PATCH")
	//r.HandleFunc("/api/v2/accounts/{id}", s.DeleteAccount).Methods("DELETE")

	// Roles routes
	r.Handle("/api/v2/roles", s.CreateRole()).Methods("POST")
	r.Handle("/api/v2/roles", s.ListRoles()).Methods("GET")
	//r.HandleFunc("/api/v2/roles/{id}", s.GetRole).Methods("GET")
	//r.HandleFunc("/api/v2/roles/{id}", s.ModifyRole).Methods("PATCH")
	//r.HandleFunc("/api/v2/roles/{id}", s.DeleteRole).Methods("DELETE")

	s.Router = r
}