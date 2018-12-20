package cognitive

import "errors"

var (
	ErrIncorrectStatusCode = errors.New("incorrect status code")
	ErrNotFound            = errors.New("not found")
)
