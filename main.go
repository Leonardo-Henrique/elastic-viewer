package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Leonardo-Henrique/elastic-viewer/src/api/routes"
	_ "github.com/Leonardo-Henrique/elastic-viewer/src/config"
)

func main() {

	r := routes.GenerateRouter()

	log.Println("Elastic Viewer Just Started without any errors")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":3000"), r))

}
