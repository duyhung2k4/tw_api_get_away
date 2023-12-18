package config

import (
	"net/http/httputil"
	"net/url"
	"tw_api_get_away/constant"
)

var (
	urlAccountService *url.URL
	rp                *httputil.ReverseProxy
	mapURL            map[constant.URL_SERVICE]*url.URL
)

const (
	URL_ACCOUNT_SERVICE = "URL_ACCOUNT_SERVICE"
)
