package server

import (
	"fmt"
	"net/http"

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

		values := map[string]string{}
		for k, v := range r.URL.Query() {
			values[k] = v[0]
		}

		defer r.Body.Close()
		route.Handler.Process(w, r.Body, values)

	})
	return http.ListenAndServe(config.ServerAddress, nil)
}
