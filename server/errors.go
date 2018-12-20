package server

import "errors"

var (
	ErrEmptyPhone      = errors.New("empty phone")
	ErrInvalidPhone    = errors.New("invalid phone")
	ErrEmptyImage      = errors.New("empty image")
	ErrEmptyParameter  = errors.New("empty parameter")
	ErrAliasNotDefined = errors.New("alias not defined")
)
