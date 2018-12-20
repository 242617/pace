package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/242617/pace/config"
)

const (
	HeaderCookie = "X-Cookie"
	HeaderToken  = "X-Token"
)

func Init() error {
	fmt.Printf("server started at %s\n", config.ServerAddress)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Cache-Control, Authorization, Origin, Content-Type, RequestToken, X-Token, X-Cookie")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == http.MethodOptions {
			return
		}

		if strings.HasPrefix(r.RequestURI, "/images") {
			name := strings.TrimPrefix(r.RequestURI, "/images")
			log.Println("name", name)
			filename := config.ImagesPath + name
			log.Println("filename", filename)

			file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
			defer file.Close()

			if err != nil {
				log.Println("err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = io.Copy(w, file)
			if err != nil {
				log.Println("err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		}

		ok, name, route := Routes.Get(r.RequestURI, r.Method)
		if !ok {
			http.Error(w, "not implemented", http.StatusNotImplemented)
			return
		}
		log.Println("name", name)

		parameters := map[string]string{}
		for k, v := range r.URL.Query() {
			parameters[k] = v[0]
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		defer r.Body.Close()
		params := route.Handler.Parameters()
		if params != nil {
			err := params.Apply(ctx, parameters, r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		headers := map[string]string{}
		for k, v := range r.Header {
			headers[k] = v[0]
		}

		route.Handler.Process(ctx, w, headers, params)

	})
	return http.ListenAndServe(config.ServerAddress, nil)
}
