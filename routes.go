package main

import (
	"net/http"
	"strings"

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
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	{
		"CreateDoc",
		strings.ToUpper("Post"),
		"/api/v1/docs",
		CreateDoc,
	},

	{
		"DeleteDoc",
		strings.ToUpper("Delete"),
		"/api/v1/docs/{id}/{version}",
		DeleteDoc,
	},

	{
		"GetDoc",
		strings.ToUpper("Get"),
		"/api/v1/docs/{id}/{version}",
		GetDoc,
	},

	{
		"GetDocs",
		strings.ToUpper("Get"),
		"/api/v1/docs",
		GetDocs,
	},
}
