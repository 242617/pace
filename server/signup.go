package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/242617/pace/storage"
)

type signup struct {
	Phone string `json:"phone"`
}

func (*signup) Parameters() parameters { return &signup{} }
func (*signup) Process(ctx context.Context, w http.ResponseWriter, parameters parameters) {
	params := parameters.(*signup)

	user, err := storage.GetUser(ctx, params.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user.Name)

}
