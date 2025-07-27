package repository

import "errors"

var (
	ErrUserExists       = errors.New("user alredy exists")
	ErrUserDoesntExists = errors.New("user doens't exists")
	ErrDBUnavailable    = errors.New("Database temporarily unavailable")
)
