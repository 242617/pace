package server

import (
	"fmt"
	"net/http"
)

type healthcheck struct{ empty }

func (*healthcheck) Process(w http.ResponseWriter, parameters parameters) {
	fmt.Fprintf(w, "status ok")
}
