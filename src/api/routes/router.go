package routes

import (
	"github.com/Leonardo-Henrique/elastic-viewer/src/api/routes/routes"
	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	return routes.InjectRoutes(mux.NewRouter())
}
