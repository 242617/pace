package server

import (
	"context"
	"encoding/json"
	"io"
)

func (a *alias) Apply(ctx context.Context, parameters map[string]string, body io.Reader) error {

	err := json.NewDecoder(body).Decode(a)
	if err != nil {
		return err
	}
	if a.Image == "" {
		return ErrEmptyImage
	}

	return nil
}
