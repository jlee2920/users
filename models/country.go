package models

// Country customizes required titles for names and
type Country struct {
	Name                string
	ShortName           string
	RequiredTitles      RequiredTitles
	RequiredAddressComp RequiredAddressComp
	RequiredTaxID       string // ex: SSN
}

// RequiredTitles refers to the required fields in a Name which will be decided on a country basis
type RequiredTitles struct {
	Titles []string
}

// RequiredAddressComp is a list of fields required in a country to define an address
type RequiredAddressComp struct {
	AddressComponents []string
}
