package cognitive

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/242617/pace/config"
)

func Identify(faceID string) (string, float64, error) {

	buf := bytes.NewBuffer([]byte{})
	request := struct {
		FaceIds                    []string `json:"faceIds"`
		PersonGroupId              string   `json:"personGroupId"`
		MaxNumOfCandidatesReturned int      `json:"maxNumOfCandidatesReturned"`
		confidenceThreshold        float64  `json:"confidenceThreshold"`
	}{[]string{faceID}, groupID, 1, .5}
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		log.Println("err", err)
		return "", 0, err
	}

	req, err := http.NewRequest(http.MethodPost, config.CognitiveURL+"/identify", buf)
	if err != nil {
		log.Println("err", err)
		return "", 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", config.CognitiveKey)

	client := &http.Client{Timeout: DefaultTimeout}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return "", 0, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("res.StatusCode", res.StatusCode)
		return "", 0, ErrIncorrectStatusCode
	}

	var response []struct {
		FaceID     string `json:"faceId"`
		Candidates []struct {
			PersonId   string  `json:"personId"`
			Confidence float64 `json:"confidence"`
		} `json:"candidates"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("err", err)
		return "", 0, err
	}

	if len(response) == 0 || len(response[0].Candidates) == 0 {
		return "", 0, ErrNotFound
	}

	personID := response[0].Candidates[0].PersonId
	confidence := response[0].Candidates[0].Confidence

	return personID, confidence, nil
}
