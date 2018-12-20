package piggybox

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SMSRequest(rawPhone string) (string, string, error) {

	buf := bytes.NewBuffer([]byte{})

	phone, err := strconv.ParseInt(rawPhone, 10, 0)
	if err != nil {
		return "", "", err
	}

	request := struct {
		Phone int64 `json:"phone"`
	}{phone}
	err = json.NewEncoder(buf).Encode(request)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	req, err := http.NewRequest(http.MethodPost, url+"/sms/request", buf)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	client := http.Client{Timeout: time.Minute}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("res.StatusCode", res.StatusCode)
		return "", "", ErrIncorrectStatusCode
	}

	var cookie string
	for _, v := range res.Cookies() {
		if v.Name == "piggybox-session" {
			cookie = v.String()
			break
		}
	}

	var response struct {
		Data struct {
			Code string `json:"code"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	return response.Data.Code, cookie, nil

}
