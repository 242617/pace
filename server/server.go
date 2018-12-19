package server

import (
	"fmt"
	"net/http"

	"github.com/242617/pace/config"
)

func Init() error {
	fmt.Printf("server started at %s\n", config.ServerAddress)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
		fmt.Println("ok")
	})
	return http.ListenAndServe(config.ServerAddress, nil)
}
