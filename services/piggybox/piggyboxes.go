package piggybox

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/242617/pace/model"
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
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode", res.StatusCode)
		defer res.Body.Close()
		barr, _ := ioutil.ReadAll(res.Body)
		log.Println(string(barr))
		return nil, ErrIncorrectStatusCode
	}

	var response struct {
		Data []struct {
			Alias  string  `json:"alias"`
			Name   string  `json:"name"`
			Amount float64 `json:"recommended_amount"`
			Status int     `json:"status"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	var piggyboxes []model.Piggybox
	for _, piggybox := range response.Data {
		if piggybox.Status == StatusActive {
			piggyboxes = append(piggyboxes, model.Piggybox{
				Alias:  piggybox.Alias,
				Name:   piggybox.Name,
				Amount: piggybox.Amount,
			})
		}
	}

	return piggyboxes, nil
}
