package models

import "github.com/pborman/uuid"

// User will be a user of the WON platform
type User struct {
	Person PhysicalPerson		// This is going under the assumption that a PhysicalPerson can have multiple users/accounts in our system
	UUID uuid.UUID				// Do we need to add this ourselves or is there some mechanism that will automatically add a struct in?
	Blockchain []Blockchain
	FlaggedBehavior []FlaggedBehavior	// This will be used when a user flagged
	Status []AccountStatus
	// TODO: Think about stuff like createdAt/modifiedAt/deletedAt - SHOULD be added by whatever ORM we are using if we're using a relational DB
	// TODO: Are we going to make a username/password schema or stick with the current Phone SMS check?
}

// AccountStatus will determine the status of the account
type AccountStatus struct { // ex: Active, Frozen, Suspended, Inactive?, etc
	StatusName string
	CanAccess bool 				// This will be a single bool that will control if a user can access this account - might want to elaborate/think on
}

// FlaggedBehavior is struct that defines specific negative behaviors about the User
type FlaggedBehavior struct { // ex: Dangerous, Risky, etc
	BehaviorName string
	BehaviorSeverity int		// BehaviorSeverity will determine the severity of the behavior
}

// Blockchain will need to contain the address on the WON blockchain
type Blockchain struct {
	Address string
	PublicKey string // TODO: Change this type when we determine how we want to make our key pairs
	// TODO: What else do we need to put in this struct
}

/* ******************************************************************************************************************************* */

// Group will contain a list of users and a list of roles that it will fulfill which is determined by permissions
type Group struct {
	UUID uuid.UUID  	// Need some way of uniquely identifying the group - we can allow uniqueness via Group name if we want
	GroupName string
	Users []User
	Roles []Role
}

// Role is a struct that groups a list of Permissions that would be assigned to a group
type Role struct {
	RoleName string
	Permissions []Permission
}

// Permission is a struct that will determine what a user can do/access
type Permission struct {
	PermissionName string
	Actions []string // Have some sort of way to create a permission that will have a list of actions it can do/what it can access
}