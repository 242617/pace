package cognitive

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/242617/pace/config"
)

func Detect(reader io.Reader) (string, error) {

	req, err := http.NewRequest(http.MethodPost, config.CognitiveURL+"/detect", reader)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Ocp-Apim-Subscription-Key", config.CognitiveKey)

	client := &http.Client{Timeout: DefaultTimeout}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode", res.StatusCode)

		defer res.Body.Close()
		barr, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(barr))

		return "", ErrIncorrectStatusCode
	}

	var response []struct {
		FaceID string `json:"faceId"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", err
	}

	return response[0].FaceID, nil

}
