package controller

import (
	"context"
	"users.git/models"
	"users.git/storage"
)

// CreateUser takes in a database and a users post struct and returns the added structure
func Ping(ctx context.Context) (*models.Ping, error) {
	// Call the database with a defined function in the server
	z := storage.Storage{}
	ping, _ := ping(ctx, z)

	return ping, nil
}

// createUser is a helper function that will take something that satisfies the usersService interface and the user object
func ping(ctx context.Context, healthCheck storage.HealthCheck) (*models.Ping, error) {
	return healthCheck.Ping(ctx)
}