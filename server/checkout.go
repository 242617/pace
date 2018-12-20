package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/242617/pace/services/cognitive"
	"github.com/242617/pace/storage"
)

type checkout struct {
	Image string `json:"image"`
}

func (*checkout) Parameters() parameters { return &checkout{} }
func (*checkout) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*checkout)
	phone := headers["Phone"]
	log.Println("phone", phone)

	params.Image = params.Image[strings.IndexByte(params.Image, ',')+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(params.Image))

	faceID, err := cognitive.Detect(reader)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("faceID", faceID)

	personID, _, err := cognitive.Identify(faceID)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("personID", personID)

	// name, data, err := cognitive.Person(personID)
	// if err != nil {
	// 	log.Println("err", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// log.Println(name, data)

	receiver, err := storage.GetUserByPersonID(ctx, personID)
	if err != nil {
		log.Println("err", err)
		switch err {
		case storage.ErrNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	log.Println(receiver)
	log.Println("receiver.Alias", receiver.Alias)

	response := struct {
		Alias string `json:"alias"`
	}{receiver.Alias}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
