package server

import (
	"fmt"
	"io"
	"net/http"
)

type healthcheck struct{ empty }

func (*healthcheck) Process(w http.ResponseWriter, b io.Reader, v map[string]string) {
	fmt.Fprintf(w, "status ok")
}
