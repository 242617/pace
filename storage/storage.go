package storage

import "errors"

var (
	ErrNotImplemented = errors.New("not implemented")
)

func Init() error {
	return nil
}
