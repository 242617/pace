package cognitive

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/242617/pace/config"
)

func AddFace(personID string, reader io.Reader) (string, error) {

	req, err := http.NewRequest(http.MethodPost, config.CognitiveURL+"/persongroups/"+groupID+"/persons/"+personID+"/persistedFaces", reader)
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
		barr, _ := ioutil.ReadAll(res.Body)
		log.Println(string(barr))
		return "", ErrIncorrectStatusCode
	}

	var response struct {
		PersistedFaceID string `json:"persistedFaceId"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", err
	}

	return response.PersistedFaceID, nil

}
