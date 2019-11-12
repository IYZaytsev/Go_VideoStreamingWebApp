package main

import (
	"net/http"

	"./handlefunc"
)

//Route Used to match requets with approaite handlers
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice holds all information on routes and what handler function to use
type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		handlefunc.Index,
	},
	Route{
		"index",
		"GET",
		"/upload",
		handlefunc.Upload,
	},
}
