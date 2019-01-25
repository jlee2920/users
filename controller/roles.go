package controller

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	httpModel "github.com/jlee2920/cosmo/structs"
	"users.git/models"
	"users.git/storage"
)

// CreateAccount will take the given format of the account and create an account on the WON platform
func CreateRole(ctx context.Context, db *gorm.DB, r *httpModel.RolePost) (*models.Role, error) {
	// Ensure we have all the required fields in place
	if r.Rolename == "" || len(r.Permissions) == 0 {
		return nil, errors.New("no rolename or permissions")
	}

	// Call the database with a defined function in the server
	z := storage.Storage{DB: db}
	role, err := createRole(ctx, z, r)
	if err != nil {
		// Return soAlsomething went wrong with database transaction
	}

	return role, nil
}

// createRole is a helper function that will take something that satisfies the usersService interface and the user object
func createRole(ctx context.Context, rolesService storage.RolesService, u *httpModel.RolePost) (*models.Role, error) {
	return rolesService.PostRole(ctx, u)
}


func ListRoles(ctx context.Context, db *gorm.DB, l *httpModel.ListModel) (*httpModel.ListModel, error) {
	z := storage.Storage{DB: db}
	list, err := listRoles(ctx, z, l)
	if err != nil {
		// Return something went wrong with database transaction
	}
	return list, nil
}

// listUsers is a helper function that will take in something that satisfies the usersService interface and the list of filters
func listRoles(ctx context.Context, rolesService storage.RolesService, l *httpModel.ListModel) (*httpModel.ListModel, error) {
	return rolesService.ListRoles(ctx, l)
}