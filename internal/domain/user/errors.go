package user

import "errors"

var (
	ErrInvalidRole     = errors.New("invalid role")
	ErrEmptyName       = errors.New("empty name")
	ErrDuplicateEmail  = errors.New("email already in use")
	ErrAddressNotFound = errors.New("address not found")
	ErrEmptyPassword   = errors.New("empty password")
)
