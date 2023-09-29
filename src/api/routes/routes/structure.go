package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouteStructure struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func InjectRoutes(r *mux.Router) *mux.Router {
	var routes []RouteStructure
	routes = append(routes, ClusterRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
