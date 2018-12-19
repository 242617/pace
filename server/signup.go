package server

import (
	"fmt"
	"net/http"
)

type signup struct {
	Phone int64 `json:"phone"`
}

func (s *signup) Parameters() parameters { return &signup{} }
func (*signup) Process(w http.ResponseWriter, parameters parameters) {
	params := parameters.(*signup)

	fmt.Println("params.Phone", params.Phone)

	// user, err := storage.GetUser(s.Phone)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(user.Name)

}
