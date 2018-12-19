package server

import (
	"net/http"
)

type route struct {
	URL     string
	Method  string
	Handler handler
}
type routes map[string]route

func (r routes) Get(url, method string) (bool, string, *route) {
	for k, v := range r {
		if v.URL == url && v.Method == method {
			return true, k, &v
		}
	}
	return false, "", nil
}

var Routes = routes{
	"healthcheck": {"/healthcheck", http.MethodGet, &healthcheck{}},
	"signup":      {"/signup", http.MethodPost, &signup{}},
}
