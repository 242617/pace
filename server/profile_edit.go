package server

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/242617/pace/storage"

	"github.com/242617/pace/services/cognitive"
)

type profile_edit struct {
	Name  *string `json:"name"`
	Alias *string `json:"alias"`
	Image *string `json:"image"`
}

func (*profile_edit) Parameters() parameters { return &profile_edit{} }
func (*profile_edit) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*profile_edit)
	phone := headers["Phone"]
	log.Println("phone", phone)

	user, err := storage.GetUserByPhone(ctx, phone)
	if err != nil {
		log.Println("err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(user)

	if params.Name != nil {
		err = storage.UpdateUserName(ctx, phone, *params.Name)
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if params.Alias != nil {
		err = storage.UpdateUserAlias(ctx, phone, *params.Alias)
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if params.Image != nil {

		personID, err := cognitive.CreatePerson(user.Phone, "")
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		image := *params.Image
		image = image[strings.IndexByte(image, ',')+1:]
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image))

		persistedFaceId, err := cognitive.AddFace(personID, reader)
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("persistedFaceId", persistedFaceId)

		err = cognitive.Train()
		if err != nil {
			log.Println("err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.PersonID != personID {
			err = storage.UpdateUserPersonID(ctx, phone, personID)
			if err != nil {
				log.Println("err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}

	w.WriteHeader(http.StatusAccepted)

}
