package returns

import (
	"encoding/json"
	"log"
	"net/http"
)

func StructureJsonResponse(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
