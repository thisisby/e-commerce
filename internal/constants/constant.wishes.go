package constants

import "errors"

var (
	ErrWishNotFound = errors.New("wish not found")
	ErrWishExists   = errors.New("wish already exists")
)
