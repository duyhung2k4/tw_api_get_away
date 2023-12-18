package handle_server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"tw_api_get_away/config"
)

func HandleServer() http.Handler {
	mapURL := config.GetmapURL()
	rp := config.GetRP()

	s := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var URL *url.URL
		cutURL := strings.Split(r.URL.String(), "/")
		checkCount := 0

		for _, u := range cutURL {
			for key, value := range mapURL {
				if u == string(key) {
					URL = value
					checkCount++
					break
				}
			}
			if checkCount > 0 {
				break
			}
		}

		rp = httputil.NewSingleHostReverseProxy(URL)
		rp.ServeHTTP(w, r)

	})

	return s
}
