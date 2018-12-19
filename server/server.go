package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/242617/pace/config"
)

func Init() error {
	fmt.Printf("server started at %s\n", config.ServerAddress)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ok, name, route := Routes.Get(r.RequestURI, r.Method)
		if !ok {
			http.Error(w, "not implemented", http.StatusNotImplemented)
			return
		}
		fmt.Println("name", name)

		parameters := map[string]string{}
		for k, v := range r.URL.Query() {
			parameters[k] = v[0]
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		defer r.Body.Close()
		params := route.Handler.Parameters()
		err := params.Apply(ctx, parameters, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		route.Handler.Process(w, params)

	})
	return http.ListenAndServe(config.ServerAddress, nil)
}
