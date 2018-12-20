package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/242617/pace/model"
	"github.com/242617/pace/services/piggybox"
	"github.com/242617/pace/storage"
)

type profile_get struct{ empty }

func (*profile_get) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	phone := headers["Phone"]
	log.Println("phone", phone)

	cookie, token := headers[HeaderCookie], headers[HeaderToken]
	log.Println("cookie", cookie)
	log.Println("token", token)

	user, err := storage.GetUserByPhone(ctx, phone)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	piggyboxes, err := piggybox.Piggyboxes(token, cookie)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("piggyboxes", piggyboxes)
	response := struct {
		Name       string           `json:"name"`
		Alias      string           `json:"alias"`
		Piggyboxes []model.Piggybox `json:"piggyboxes"`
	}{user.Name, user.Alias, piggyboxes}

	w.WriteHeader(http.StatusAccepted)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
