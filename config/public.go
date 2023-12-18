package config

import (
	"net/http/httputil"
	"net/url"
	"tw_api_get_away/constant"
)

func GetURLAccountService() *url.URL {
	return urlAccountService
}

func GetRP() *httputil.ReverseProxy {
	return rp
}

func GetmapURL() map[constant.URL_SERVICE]*url.URL {
	mapURL := map[constant.URL_SERVICE]*url.URL{
		constant.ACCOUNT_SERVICE: urlAccountService,
	}
	return mapURL
}
