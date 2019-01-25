package storage

import (
	"context"
	httpModel "github.com/jlee2920/cosmo/structs"
	"users.git/models"
)

// Example just for testing
var exampleRole = &models.Role{RoleName: "example", Permissions: []models.Permission{examplePermission}}
var examplePermission = models.Permission{PermissionName: "admin", Actions: []string{"all"}}

func (s Storage) PostRole(ctx context.Context, r *httpModel.RolePost) (*models.Role, error) {
	// TODO: Conduct database transaction
	return exampleRole, nil
}
func (s Storage) GetRole(ctx context.Context, r string) (*models.Role, error) {
	return nil, nil
}
func (s Storage) PatchRole(ctx context.Context, r *models.Role) (*models.Role, error) {
	return nil, nil
}
func (s Storage) ListRoles(ctx context.Context, l *httpModel.ListModel) (*httpModel.ListModel, error) {
	// TODO: Conduct database transaction
	newModel := &httpModel.ListModel{
		Skip: 0,
		Limit: 10,
		Resources: []string{exampleRole.RoleName},
	}
	return newModel, nil
}