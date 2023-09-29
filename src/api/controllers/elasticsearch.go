package controllers

import (
	"log"
	"net/http"

	"github.com/Leonardo-Henrique/elastic-viewer/src/api/returns"
	"github.com/Leonardo-Henrique/elastic-viewer/src/api/services"
)

func ReturnClusterStatus(w http.ResponseWriter, r *http.Request) {
	ec := services.MakeESConn()
	status, err := ec.GetClusterStats()
	if err != nil {
		log.Fatal(err)
	}
	returns.StructureJsonResponse(w, http.StatusOK, status)
}
