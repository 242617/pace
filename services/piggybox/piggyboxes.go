package piggybox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.qiwi.com/ssk-dev-apps/piggybox-telegram-bot/model"
)

func Piggyboxes(token, cookie string) ([]model.Piggybox, error) {

	req, err := http.NewRequest(http.MethodGet, url+"/piggyboxes", nil)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RequestToken", token)
	req.Header.Set("Cookie", cookie)

	client := http.Client{Timeout: time.Minute}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode", res.StatusCode)
		defer res.Body.Close()
		barr, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(barr))
		return nil, ErrIncorrectStatusCode
	}

	var response struct {
		Data []struct {
			Alias  string  `json:"alias"`
			Name   string  `json:"name"`
			Amount float64 `json:"recommended_amount"`
		} `json:"data"`
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	fmt.Println(response.Data)

	return nil, nil
}
