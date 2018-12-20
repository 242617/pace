package piggybox

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Create(token string, cookie string) (string, error) {

	req, err := http.NewRequest(http.MethodPost, url+"/piggybox", strings.NewReader("{}"))
	if err != nil {
		log.Println("err", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RequestToken", token)
	req.Header.Set("Cookie", cookie)

	client := http.Client{Timeout: time.Minute}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("res.StatusCode", res.StatusCode)
		defer res.Body.Close()
		barr, _ := ioutil.ReadAll(res.Body)
		log.Println(string(barr))
		return "", ErrIncorrectStatusCode
	}

	var response struct {
		Data struct {
			Alias string `json:"alias"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", err
	}

	alias := response.Data.Alias
	log.Println("alias", alias)

	return alias, nil

}
