package constants

import (
	"errors"
)

var (
	ErrRowExists   = errors.New("row already exists")
	ErrRowNotFound = errors.New("row not found")
)
