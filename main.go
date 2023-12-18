package main

import (
	"net/http"
	"time"
	handle_server "tw_api_get_away/handleServer"
)

func main() {
	s := http.Server{
		Addr:         ":18888",
		Handler:      handle_server.HandleServer(),
		ReadTimeout:  6000 * time.Second,
		WriteTimeout: 6000 * time.Second,
	}

	s.ListenAndServe()
}
