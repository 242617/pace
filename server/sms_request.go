package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/242617/pace/services/piggybox"
)

type sms_request struct {
	Phone string `json:"phone"`
}

func (*sms_request) Parameters() parameters { return &sms_request{} }
func (*sms_request) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*sms_request)

	code, cookie, err := piggybox.SMSRequest(params.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("code", code)
	log.Println("cookie", cookie)

	response := struct {
		Code   string `json:"code"`
		Cookie string `json:"cookie"`
	}{code, cookie}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
