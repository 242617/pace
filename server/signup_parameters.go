package server

import (
	"context"
	"encoding/json"
	"io"
)

func (s *signup) Apply(ctx context.Context, parameters map[string]string, body io.Reader) error {

	err := json.NewDecoder(body).Decode(s)
	if err != nil {
		return err
	}

	return nil
}
