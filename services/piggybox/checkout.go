package piggybox

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Checkout(alias string, amount float64, rawPhone string) (string, string, error) {

	phone, err := strconv.ParseInt(rawPhone, 10, 0)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	buf := bytes.NewBuffer([]byte{})
	request := struct {
		Amount     float64 `json:"amount"`
		CurrencyID int     `json:"currency_id"`
		Source     string  `json:"source"`
		Payer      int64   `json:"payer"`
		Comment    string  `json:"comment"`
	}{amount, 643, "card", phone, "pace payment"}
	err = json.NewEncoder(buf).Encode(request)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	req, err := http.NewRequest(http.MethodPost, url+"/checkout/"+alias, buf)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: time.Minute}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("res.StatusCode", res.StatusCode)
		barr, _ := ioutil.ReadAll(res.Body)
		log.Println(string(barr))
		return "", "", ErrIncorrectStatusCode
	}

	var response struct {
		Data struct {
			URL         string `json:"url"`
			Transaction string `json:"transaction"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	transaction := response.Data.Transaction
	url := response.Data.URL

	return transaction, url, nil

}
