package main

import (
	"log"
	"net/http"
	"users.git/config"
	"users.git/middleware"
	"users.git/routes"
)

var server routes.Server

func init() {
	server.Config = config.InitializeConfig()
	//server.DB = storage.InitializeDB(server.Config)
	server.InitializeRoutes()
}

func main() {
	// Apply general middleware here while route specific middleware will be applied during InitializeRoutes
	//var app = middleware.Middleware(server.Router, middleware.AuthMiddleware, middleware.RequestIDMiddlware)

	var app = middleware.Middleware(server.Router, middleware.RequestIDMiddlware)
	// Run the app
	log.Fatal(http.ListenAndServe(":8080", app))
}
