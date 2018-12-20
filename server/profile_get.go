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
	log.Println("len(piggyboxes)", len(piggyboxes))

	response := struct {
		Name       string           `json:"name"`
		Alias      string           `json:"alias"`
		Piggyboxes []model.Piggybox `json:"piggyboxes"`
	}{}
	response.Name = user.Name
	response.Alias = user.Alias

	if len(piggyboxes) == 0 {

		alias, err := piggybox.Create(token, cookie)
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("alias", alias)

		piggyboxes, err = piggybox.Piggyboxes(token, cookie)
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("len(piggyboxes)", len(piggyboxes))

		response.Alias = ""
		err = storage.UpdateUserAlias(ctx, phone, "")
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	response.Piggyboxes = piggyboxes

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
