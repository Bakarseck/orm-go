package server

import "net/http"

type EndPoint struct {
	Path    string
	Handler RouteHandler
	Method  string
}

type RouteHandler func(http.ResponseWriter, *http.Request)

type Middleware func(RouteHandler) RouteHandler

type route struct {
	method  string
	pattern string
	handler RouteHandler
}

type Router struct {
	routes      []route
	middlewares []Middleware
}

var (
	Endpoints = []EndPoint{}
)
