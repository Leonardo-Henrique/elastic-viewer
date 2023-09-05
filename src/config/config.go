package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ELASTICSEARCH_HOST     string
	ELASTICSEARCH_USERNAME string
	ELASTICSEARCH_PASSWORD string
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file detected")
	}

	if eH := os.Getenv("ELASTICSEARCH_HOST"); eH != "" {
		ELASTICSEARCH_HOST = eH
	}

	if eU := os.Getenv("ELASTICSEARCH_USERNAME"); eU != "" {
		ELASTICSEARCH_USERNAME = eU
	}

	if eP := os.Getenv("ELASTICSEARCH_PASSWORD"); eP != "" {
		ELASTICSEARCH_PASSWORD = eP
	}

}
