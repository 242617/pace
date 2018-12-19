package server

import "errors"

var (
	ErrEmptyPhone   = errors.New("empty phone")
	ErrInvalidPhone = errors.New("invalid phone")
)
