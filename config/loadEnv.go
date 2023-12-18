package config

import (
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	urlAccountServiceConvert, errAccountService := url.Parse(os.Getenv(URL_ACCOUNT_SERVICE))
	if errAccountService != nil {
		return errAccountService
	}
	urlAccountService = urlAccountServiceConvert

	return nil
}
