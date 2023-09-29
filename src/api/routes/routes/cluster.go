package routes

import (
	"net/http"

	"github.com/Leonardo-Henrique/elastic-viewer/src/api/controllers"
)

var ClusterRoutes = []RouteStructure{
	{
		URI:    "/cluster",
		Method: http.MethodGet,
		Func:   controllers.ReturnClusterStatus,
	},
}
