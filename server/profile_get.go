package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/242617/pace/model"
	"github.com/242617/pace/services/piggybox"
)

type profile_get struct{ empty }

func (*profile_get) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {

	cookie, token := headers[HeaderCookie], headers[HeaderToken]
	fmt.Println("cookie", cookie)
	fmt.Println("token", token)

	piggyboxes, err := piggybox.Piggyboxes(token, cookie)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("piggyboxes", piggyboxes)
	response := struct {
		Name       string           `json:"name"`
		Piggyboxes []model.Piggybox `json:"piggyboxes"`
	}{"user.Name", piggyboxes}

	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
