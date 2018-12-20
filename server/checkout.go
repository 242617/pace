package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/242617/pace/config"

	"github.com/242617/pace/services/piggybox"
)

type checkout struct {
	Alias  string  `json:"alias"`
	Amount float64 `json:"amount"`
}

func (*checkout) Parameters() parameters { return &checkout{} }
func (*checkout) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*checkout)
	phone := headers["Phone"]
	log.Println("phone", phone)

	log.Println("params.Alias", params.Alias)
	log.Println("params.Amount", params.Amount)

	transaction, url, err := piggybox.Checkout(params.Alias, params.Amount, phone)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url, err = fix(url, config.SuccessPage)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Transaction string `json:"transaction"`
		Url         string `json:"url"`
	}{transaction, url}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func fix(raw string, successUrl string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("successUrl", successUrl)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
