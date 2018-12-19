package server

import (
	"encoding/json"
	"io"
)

func (s *signup) Validate(values map[string]string, body io.Reader) error {

	err := json.NewDecoder(body).Decode(s)
	if err != nil {
		return err
	}

	return nil
}
