package server

import (
	"context"
	"io"
	"net/http"
)

type handler interface {
	Process(context.Context, http.ResponseWriter, parameters)
	Parameters() parameters
}

type parameters interface {
	Apply(context.Context, map[string]string, io.Reader) error
}

type empty struct{}

func (*empty) Parameters() parameters                                    { return nil }
func (*empty) Apply(context.Context, map[string]string, io.Reader) error { return nil }
