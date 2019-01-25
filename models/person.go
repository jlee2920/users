package models

// PhysicalPerson is a person that uses our system - this is separate from account since a physicalperson can technically have multiple accounts
type PhysicalPerson struct {
	Name         Name
	Address      PhysicalAddress
	PhoneNumbers []string
	Emails       []string
	Citizenship  []string
	TaxID        TaxID
	DOB          string // Should we account for lunar and would that make a difference?
	Fund         FundSource
	IDDocuments  []IDDocuments
	KYCLevel     int
	PinCode      int
}

// Name defines a particular string with a portion of a name of a physical person
type Name struct {
	TitleOfName string //first name, last name, etc
	Name        string
}

// PhysicalAddress should have all the components an address composes of
type PhysicalAddress struct {
	Country Country
	Address []string // List of address strings that comply with the Country
}

// TaxID contains the required TaxID in the particular country
type TaxID struct {
	Country Country
	TaxID   string // Could be SSN or some other uniquely, government approved identifiable string
}

// FundSource is the way a user will put fiat in or take fiat out of our system
type FundSource struct {
	CreditCard []CreditCardInformation
	Bank       []BankInformation
}

// CreditCardInformation should contain the required information to make a transaction with the credit card company
type CreditCardInformation struct {
	// TODO: Figure out what we exactly need in terms of credit card information
	KYCLevel	int // We need to vet the the credit card company in order to take fiat in
}

// BankInformation would be the bank information the user should provide
type BankInformation struct {
	BankName      string
	Address       PhysicalAddress
	RoutingNumber string
	AccountNumber string
	// TODO: other needed bank information
	ClearingInformation BankClearingInformation // ex: the US Fed/Swift
	KYCLevel	int // We need to vet the the bank in order to take fiat in
}

// BankClearingInformation contains the required information needed to ensure athe legitimacy of the bank
type BankClearingInformation struct {
	SystemName  string
	NumberToUse int
}

// IDDocuments will hold any PII regarding a user
type IDDocuments struct {
	// TODO: Figure out what documents we need to keep on behalf of the individual to identify them
	Title   string
	UUID    string
	Country Country
}
