package main

import "github.com/pkg/errors"

var (
	ErrBadRequest       = errors.New("Bad Request")
	ErrForbidden        = errors.New("Forbidden")
	ErrNotFound         = errors.New("Not Found")
	ErrMethodNotAllowed = errors.New("Method Not Allowed")
)
