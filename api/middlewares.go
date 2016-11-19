package api

import (
	"fmt"
	"net/http"
	"strings"
)

func (rc *V1) AddDefaultHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Adding default header..\n")
		w.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, Accept, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if strings.Compare(r.Method, "OPTIONS") == 0 {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rc *V1) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, wl := range rc.Whitelist {
			if wl.MatchString(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		accessToken := r.Header.Get("access_token")
		if accessToken != "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rc.AccessToken = accessToken
		next.ServeHTTP(w, r)
	})
}
