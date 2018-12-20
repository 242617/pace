package server

import (
	"context"
	"net/http"
)

type profile_edit struct {
	Name  string  `json:"name"`
	Alias string  `json:"alias"`
	Image *string `json:"image"`
}

func (*profile_edit) Parameters() parameters { return &profile_edit{} }
func (*profile_edit) Process(ctx context.Context, w http.ResponseWriter, headers headers, parameters parameters) {
	params := parameters.(*profile_edit)

	if params.Image != nil {

		// image := *params.Image
		// image = image[strings.IndexByte(image, ',')+1:]

		// err := cognitive.CreatePerson(phone)
		// if err != nil {
		// 	log.Println("err", err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// faceID, err := cognitive.Detect(url)
		// if err != nil {
		// 	log.Println("err", err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// fmt.Println("faceID", faceID)

	}

	w.WriteHeader(http.StatusAccepted)

}
