package cognitive

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/242617/pace/config"
)

func Train() error {

	req, err := http.NewRequest(http.MethodPost, config.CognitiveURL+"/persongroups/"+groupID+"/train", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", config.CognitiveKey)

	client := &http.Client{Timeout: DefaultTimeout}
	res, err := client.Do(req)
	if err != nil {
		log.Println("err", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusAccepted {
		log.Println("res.StatusCode", res.StatusCode)
		barr, _ := ioutil.ReadAll(res.Body)
		log.Println(string(barr))
		return ErrIncorrectStatusCode
	}

	return nil

}
