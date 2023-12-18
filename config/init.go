package config

import "log"

func init() {
	if err := loadEnv(); err != nil {
		log.Println(err)
		return
	}
}
