package constants

import "errors"

var (
	ErrCartItemsNotFound = errors.New("cart_items not found")
	ErrCartItemsExists   = errors.New("cart_items already exists")
)
