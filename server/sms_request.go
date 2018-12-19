package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/242617/pace/services/piggybox"
)

type sms_request struct {
	Phone string `json:"phone"`
}

func (*sms_request) Parameters() parameters { return &sms_request{} }
func (*sms_request) Process(ctx context.Context, w http.ResponseWriter, parameters parameters) {
	params := parameters.(*sms_request)

	code, cookie, err := piggybox.SMSRequest(params.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Println("code", code)
	fmt.Println("cookie", cookie)

	response := struct {
		Code   string `json:"code"`
		Cookie string `json:"cookie"`
	}{code, cookie}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
