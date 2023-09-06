package main

import (
	"log"

	"github.com/Leonardo-Henrique/elastic-viewer/src/config"
	"github.com/Leonardo-Henrique/elastic-viewer/src/internal/services"
	"github.com/Leonardo-Henrique/elastic-viewer/src/models"
)

func main() {

	ec := services.MakeESConn(models.ESCredentials{
		ESHost:     config.ELASTICSEARCH_HOST,
		ESUsername: config.ELASTICSEARCH_USERNAME,
		ESPassword: config.ELASTICSEARCH_PASSWORD,
	})

	if err := ec.PingESConn(); err != nil {
		log.Fatal(err)
	}

	cluster, err := ec.GetClusterStats()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", cluster)

}
