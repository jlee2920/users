package storage

import (
	"context"
	"github.com/jinzhu/gorm"
	httpModel "github.com/jlee2920/cosmo/structs"
	"users.git/models"
)

/*	Server is the same type as routes.Server.
	Since we cannot extend existing types in another package so we defined a local alias that is a general struct
	which we can use to satisfy the UsersService interface - edit: cannot do this because of import cycle so we're just
	going to change to a new struct called Storage which will only have the db
*/
type Storage struct {
	DB *gorm.DB
}

// UsersStorage allows access to the system storing user records.
type UsersService interface {
	// Stores a user
	PostUser(ctx context.Context, u *httpModel.UsersPost) (*models.User, error)
	// Retrieve a user with a given id
	GetUser(ctx context.Context, id uint64) (*models.User, error)
	//Updates a user
	PatchUser(ctx context.Context, u *models.User) (*models.User, error)
	// Retrieves a list of users given the parameters
	ListUsers(ctx context.Context, l *httpModel.ListModel) (*httpModel.ListModel, error)
}

// RolesService allows access to the system storing roles records
type RolesService interface {
	// Stores a role
	PostRole(ctx context.Context, r *httpModel.RolePost) (*models.Role, error)
	// Retreive a role with a given name
	GetRole(ctx context.Context, r string) (*models.Role, error)
	// Updates a role
	PatchRole(ctx context.Context, r *models.Role) (*models.Role, error)
	// Retrieves a list of roles given the parameter
	ListRoles(ctx context.Context, l *httpModel.ListModel) (*httpModel.ListModel, error)
}

type HealthCheck interface {
	Ping(ctx context.Context) (*models.Ping, error)
}