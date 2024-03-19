package utils

import "errors"

var (
	ErrDataNotFound         = errors.New("data nof found")
	ErrLoginFail            = errors.New("username or password invalid")
	ErrEmailIsAlreadyExists = errors.New("this email is already exists")
)
