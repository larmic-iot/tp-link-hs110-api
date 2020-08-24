package api

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
