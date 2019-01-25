package storage

import (
	"context"
	httpModel "github.com/jlee2920/cosmo/structs"
	"users.git/models"
)

// Example for just testing
var exampleAccount = &models.User{Person: examplePerson, UUID: nil, Blockchain: nil, FlaggedBehavior: nil, Status: nil}
var examplePerson = models.PhysicalPerson{Name: models.Name{TitleOfName: "First name", Name: "John"}, Address: models.PhysicalAddress{Country: models.Country{ Name: "USA"}, Address: []string{"100 example street example CA, 10000"}}, PhoneNumbers: []string{"100-100-1000"}, Emails: []string{"example@example.com"}, Citizenship: []string{"USA"}, TaxID: models.TaxID{Country: models.Country{ Name: "USA"}, TaxID: "12345"}, DOB: "09/10/1995", Fund: models.FundSource{CreditCard: nil, Bank: nil}, IDDocuments: nil, KYCLevel: 1, PinCode: 1000}

func (s Storage) PostUser(ctx context.Context, u *httpModel.UsersPost) (*models.User, error) {
	// TODO: Conduct database transaction
	return exampleAccount, nil
}
func (s Storage) GetUser(ctx context.Context, id uint64) (*models.User, error) {
	return nil, nil
}
func (s Storage) PatchUser(ctx context.Context, u *models.User) (*models.User, error) {
	return nil, nil
}
func (s Storage) ListUsers(ctx context.Context, l *httpModel.ListModel) (*httpModel.ListModel, error) {
	// TODO: Conduct database transaction
	newModel := &httpModel.ListModel{
		Skip: 0,
		Limit: 10,
		Resources: []string{examplePerson.DOB},
	}
	return newModel, nil
}