package server

import (
	"context"
	"encoding/json"
	"io"
)

func (c *checkout) Apply(ctx context.Context, parameters map[string]string, body io.Reader) error {

	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		return err
	}
	if c.Alias == "" {
		return ErrEmptyParameter
	}
	if c.Amount == 0 {
		return ErrEmptyParameter
	}

	return nil
}
