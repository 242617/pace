package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/242617/pace/services/piggybox"
	"github.com/242617/pace/storage"
)

type sms_confirm struct {
	Code   string `json:"code"`
	VCode  string `json:"vcode"`
	Cookie string `json:"cookie"`
}

func (*sms_confirm) Parameters() parameters { return &sms_confirm{} }
func (*sms_confirm) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*sms_confirm)

	cookie := headers[HeaderCookie]

	phone, token, cookie, err := piggybox.SMSConfirm(params.Code, params.VCode, cookie)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("phone", phone)
	log.Println("token", token)
	log.Println("cookie", cookie)

	user, err := storage.GetUserByPhone(ctx, phone)
	if err == storage.ErrNotFound {

		err = storage.CreateUser(ctx, phone)
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(user)

	response := struct {
		Phone  string `json:"phone"`
		Token  string `json:"token"`
		Cookie string `json:"cookie"`
	}{phone, token, cookie}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
