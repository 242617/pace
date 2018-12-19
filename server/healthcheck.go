package server

import (
	"context"
	"fmt"
	"net/http"
)

type healthcheck struct{ empty }

func (*healthcheck) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	fmt.Fprintf(w, "status ok")
}
