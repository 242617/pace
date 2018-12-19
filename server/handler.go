package server

import (
	"io"
	"net/http"
)

type handler interface {
	Process(http.ResponseWriter, io.Reader, map[string]string)
	Validate(map[string]string, io.Reader) error
}

type empty struct{}

func (*empty) Validate(map[string]string, io.Reader) error { return nil }
