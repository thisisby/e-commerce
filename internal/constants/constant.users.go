package constants

import "errors"

const (
	Admin  = "admin"
	Client = "client"
)

var (
	// vars
	MapperRoleToId = map[string]int{
		Admin:  1,
		Client: 2,
	}
	CtxAuthenticatedUserKey = "CtxAuthenticatedUserKey"

	// errors
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)
