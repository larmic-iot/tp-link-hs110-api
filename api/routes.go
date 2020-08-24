package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = Handle404()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	Route{
		Name:        "Info",
		Method:      "GET",
		Pattern:     "/api/{ip}",
		HandlerFunc: InfoHandler,
	},
	Route{
		Name:        "Energy",
		Method:      "GET",
		Pattern:     "/api/{ip}/energy",
		HandlerFunc: EnergyHandler,
	},
}

func Handle404() http.Handler {
	type ProtocolError struct {
		Code    int
		Message string
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.
			NewEncoder(w).
			Encode(
				ProtocolError{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				})
	})
}
