package server

import (
	"context"
	"net/http"
)

type get_profile struct{ empty }

func (*get_profile) Parameters() parameters { return &get_profile{} }
func (*get_profile) Process(ctx context.Context, w http.ResponseWriter, parameters parameters) {
	// params := parameters.(*get_profile)

	// user, err := storage.GetUser(ctx, "")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// fmt.Println(user)

}
