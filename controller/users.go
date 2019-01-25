package controller

import (
	"errors"
	"github.com/jinzhu/gorm"
	httpModel "github.com/jlee2920/cosmo/structs"
	"users.git/models"
	"users.git/storage"
	"context"
)

// CreateUser takes in a database and a users post struct and returns the added structure
func CreateUser(ctx context.Context, db *gorm.DB, u *httpModel.UsersPost) (*models.User, error) {
	// Ensure we have all the required fields in place
	if u.Username == "" || u.Password == "" {
		return nil, errors.New("no username or password")
	}

	// Call the database with a defined function in the server
	z := storage.Storage{DB: db}
	account, err := createUser(ctx, z, u)
	if err != nil {
		// Return something went wrong with database transaction
	}

	return account, nil
}

// createUser is a helper function that will take something that satisfies the usersService interface and the user object
func createUser(ctx context.Context, usersService storage.UsersService, u *httpModel.UsersPost) (*models.User, error) {
	return usersService.PostUser(ctx, u)
}

// ListUsers takes in a database and a list of filters and returns the list of all users in the database
func ListUsers(ctx context.Context, db *gorm.DB, l *httpModel.ListModel) (*httpModel.ListModel, error) {
	z := storage.Storage{DB: db}
	list, err := listUsers(ctx, z, l)
	if err != nil {
		// Return something went wrong with database transaction
	}
	return list, nil
}

// listUsers is a helper function that will take in something that satisfies the usersService interface and the list of filters
func listUsers(ctx context.Context, usersService storage.UsersService, l *httpModel.ListModel) (*httpModel.ListModel, error) {
	return usersService.ListUsers(ctx, l)
}