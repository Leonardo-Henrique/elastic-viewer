package main

import (
	"log"

	"github.com/Leonardo-Henrique/elastic-viewer/src/api/services"
)

func main() {

	ec := services.MakeESConn()

	if err := ec.PingESConn(); err != nil {
		log.Fatal(err)
	}

	cluster, err := ec.GetClusterStats()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", cluster)

}
