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

type alias struct {
	Image string `json:"image"`
}

func (*alias) Parameters() parameters { return &alias{} }
func (*alias) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*alias)
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
	log.Println("receiver.Alias", receiver.Alias)
	if receiver.Alias == "" {
		http.Error(w, ErrAliasNotDefined.Error(), http.StatusBadRequest)
		return
	}

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
