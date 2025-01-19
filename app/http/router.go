package http

import (
	"fmt"
	"log"
)

// Middleware function type
type MiddlewareFunc func(req *Request, res *Response, next func())

type routeKey struct {
	Method Method
	Path   string
}

type Route struct {
	handler        func(req *Request, res *Response)
	preMiddleware  []MiddlewareFunc
	postMiddleware []MiddlewareFunc
}

type Router struct {
	routes map[routeKey]*Route
}

func NewRouter() *Router {
	router := &Router{
		routes: make(map[routeKey]*Route),
	}
	return router
}

func (r *Router) MergeRouter(other *Router) {
    for key, route := range other.routes {
        r.routes[key] = route
    }
}

func (r *Router) addRoute(method Method, path string, handler func(req *Request, res *Response)) {
	route := routeKey{Method: method, Path: path}
	r.routes[route] = &Route{
		handler:        handler,
		preMiddleware:  make([]MiddlewareFunc, 0),
		postMiddleware: make([]MiddlewareFunc, 0),
	}
}

func (r *Router) Get(path string, handler func(req *Request, res *Response)) {
	r.addRoute(GET, path, handler)
}

func (r *Router) Post(path string, handler func(req *Request, res *Response)) {
	r.addRoute(POST, path, handler)
}

func (r *Router) Put(path string, handler func(req *Request, res *Response)) {
	r.addRoute(PUT, path, handler)
}

func (r *Router) Patch(path string, handler func(req *Request, res *Response)) {
	r.addRoute(PATCH, path, handler)
}

func (r *Router) Delete(path string, handler func(req *Request, res *Response)) {
	r.addRoute(DELETE, path, handler)
}

// AddPreMiddleware adds middleware to run before the handler does run for a specific route
func (r *Router) AddPreMiddleware(method Method, path string, middleware MiddlewareFunc) error {
	key := routeKey{Method: method, Path: path}
	route, exists := r.routes[key]
	if !exists {
		return fmt.Errorf("route not found: %s %s", method, path)
	}

	route.preMiddleware = append(route.preMiddleware, middleware)
	// r.routes[key] = route
	return nil
}

// AddRouterPreMiddleware adds middleware to run before the handler does run for all routes registered
func (r *Router) AddRouterPreMiddleware(middleware MiddlewareFunc) {
	for _, route := range r.routes {
		route.preMiddleware = append(route.preMiddleware, middleware)
		// r.routes[key] = route
	}

}

// AddPostMiddleware adds middleware to run after the handler does run
func (r *Router) AddPostMiddleware(method Method, path string, middleware MiddlewareFunc) error {
	key := routeKey{Method: method, Path: path}
	route, exists := r.routes[key]
	if !exists {
		return fmt.Errorf("route not found: %s %s", method, path)
	}

	route.postMiddleware = append(route.postMiddleware, middleware)
	// r.routes[key] = route
	return nil
}

func (r *Router) AddRouterPostMiddleware(middleware MiddlewareFunc) {
	for _, route := range r.routes {
		route.postMiddleware = append(route.postMiddleware, middleware)
		// r.routes[key] = route
	}
}

func (r *Router) Resolve(req *Request, res *Response) {
	routeKey := routeKey{Method: req.GetMethod(), Path: req.GetPath()}

	route, exists := r.routes[routeKey]
	if !exists {
		res.NotFound()
		log.Print(req.GetMethod(), " ", req.GetPath(), " ", res.GetStatusCode().Int(), " ", res.GetStatusCode())
		return
	}

	// Execute pre-middleware
	for _, middleware := range route.preMiddleware {
		hasNextBeenCalled := false
		middleware(req, res, func() {
			hasNextBeenCalled = true
		})
		if !hasNextBeenCalled {
			break // Middleware chain was broken
		}
	}

	// Execute the handler
	route.handler(req, res)

	// Execute post-middleware
	for _, middleware := range route.postMiddleware {
		hasNextBeenCalled := false
		middleware(req, res, func() {
			hasNextBeenCalled = true
		})
		if !hasNextBeenCalled {
			break // Middleware chain was broken
		}
	}
}
