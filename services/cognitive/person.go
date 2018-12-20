package cognitive

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/242617/pace/config"
)

func Person(personID string) (string, string, error) {

	req, err := http.NewRequest(http.MethodGet, config.CognitiveURL+"/persongroups/"+groupID+"/persons/"+personID, nil)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", config.CognitiveKey)

	client := &http.Client{Timeout: DefaultTimeout}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode", res.StatusCode)
		return "", "", ErrIncorrectStatusCode
	}

	var response struct {
		PersonId         string   `json:"personId"`
		PersistedFaceIds []string `json:"persistedFaceIds"`
		Name             string   `json:"name"`
		Data             string   `json:"userData"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", "", err
	}

	name := response.Name
	data := response.Data

	return name, data, nil

}
