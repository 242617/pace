package piggybox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SMSConfirm(code, vcode string, cookie string) (string, string, string, error) {

	buf := bytes.NewBuffer([]byte{})
	request := struct {
		Code  string `json:"code"`
		VCode string `json:"vcode"`
	}{code, vcode}
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		log.Println("err", err)
		return "", "", "", err
	}

	req, err := http.NewRequest(http.MethodPost, url+"/sms/confirm", buf)
	if err != nil {
		log.Println("err", err)
		return "", "", "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)

	client := http.Client{Timeout: time.Minute}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		log.Println("res.StatusCode", res.StatusCode)
		return "", "", "", ErrIncorrectStatusCode
	}

	var response struct {
		Data struct {
			Phone int64  `json:"phone"`
			Token string `json:"token"`
		} `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", "", "", err
	}

	for _, v := range res.Cookies() {
		if v.Name == "piggybox-session" {
			cookie = v.String()
			break
		}
	}

	phone := strconv.FormatInt(response.Data.Phone, 10)
	fmt.Println("phone", phone)

	token := response.Data.Token
	fmt.Println("token", token)

	return phone, token, cookie, nil

}
