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
	"healthcheck":  {"/healthcheck", http.MethodGet, &healthcheck{}},
	"sms_request":  {"/sms/request", http.MethodPost, &sms_request{}},
	"sms_confirm":  {"/sms/confirm", http.MethodPost, &sms_confirm{}},
	"get_profile":  {"/profile", http.MethodGet, &profile_get{}},
	"edit_profile": {"/profile", http.MethodPut, &profile_edit{}},
	"checkout":     {"/checkout", http.MethodPost, &checkout{}},
	"alias":        {"/alias", http.MethodPost, &alias{}},
}
