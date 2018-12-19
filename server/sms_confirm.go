package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/242617/pace/services/piggybox"
)

type sms_confirm struct {
	Code   string `json:"code"`
	VCode  string `json:"vcode"`
	Cookie string `json:"cookie"`
}

func (*sms_confirm) Parameters() parameters { return &sms_confirm{} }
func (*sms_confirm) Process(ctx context.Context, w http.ResponseWriter, parameters parameters) {
	params := parameters.(*sms_confirm)

	phone, token, cookie, err := piggybox.SMSConfirm(params.Code, params.VCode, params.Cookie)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("phone", phone)
	fmt.Println("token", token)
	fmt.Println("cookie", cookie)

	response := struct {
		Phone  string `json:"phone"`
		Token  string `json:"token"`
		Cookie string `json:"cookie"`
	}{phone, token, cookie}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}