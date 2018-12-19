package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/242617/pace/storage"
)

type signup struct {
	Phone int64 `json:"phone"`
}

func (s *signup) Process(w http.ResponseWriter, b io.Reader, v map[string]string) {
	err := s.Validate(v, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("s.Phone", s.Phone)

	user, err := storage.GetUser(s.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user.Name)

}
