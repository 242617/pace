package cognitive

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/242617/pace/config"
)

func CreatePerson(name, data string) error {

	buf := bytes.NewBuffer([]byte{})
	request := struct {
		URL  string `json:"url"`
		Data string `json:"userData"`
	}{name, data}
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		log.Println("err", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, config.CognitiveURL+"/persongroups/"+groupID+"/persons", buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", config.CognitiveKey)

	client := &http.Client{Timeout: DefaultTimeout}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode", res.StatusCode)
		return ErrIncorrectStatusCode
	}

	var response []struct {
		PersonID string `json:"personId"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return err
	}

	return nil

}
