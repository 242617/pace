package server

import (
	"fmt"
	"net/http"
)

type healthcheck struct{}

func (*healthcheck) Process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status ok")
}
