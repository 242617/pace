package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/242617/pace/services/cognitive"
)

type checkout struct {
	Image string `json:"image"`
}

func (*checkout) Parameters() parameters { return &checkout{} }
func (*checkout) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*checkout)

	params.Image = params.Image[strings.IndexByte(params.Image, ',')+1:]
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(params.Image))

	faceID, err := cognitive.Detect(reader)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("faceID", faceID)

	personID, confidence, err := cognitive.Identify(faceID)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("personID", personID)
	if confidence < .6 {
		http.Error(w, "low confidence", http.StatusBadRequest)
		return
	}

	// name, data, err := cognitive.Person(personID)
	// if err != nil {
	// 	log.Println("err", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println(name, data)

	w.WriteHeader(http.StatusCreated)

}
