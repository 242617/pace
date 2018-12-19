package server

import (
	"context"
	"fmt"
	"net/http"
)

type profile_edit struct {
	Name string `json:"name"`
}

func (*profile_edit) Parameters() parameters { return &profile_edit{} }
func (*profile_edit) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*profile_edit)

	fmt.Println("params.Name", params.Name)

}
