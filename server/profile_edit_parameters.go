package server

import (
	"context"
	"encoding/json"
	"io"
)

func (p *profile_edit) Apply(ctx context.Context, parameters map[string]string, body io.Reader) error {

	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	return nil
}
