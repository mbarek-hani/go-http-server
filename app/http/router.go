package http

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

// UsePreMiddlewares adds one or more middlewares to run before the handler does run for a specific route
func (route *Route) UsePreMiddlewares(middlewares []MiddlewareFunc) *Route{
	route.preMiddleware = append(route.preMiddleware, middlewares...)
	return route
}

// UsePostMiddlewares adds one or more middlewares to run before the handler does run for a specific route
func (route *Route) UsePostMiddlewares(middlewares []MiddlewareFunc) *Route{
	route.postMiddleware = append(route.postMiddleware, middlewares...)
	return route
}

type Router struct {
	routes map[routeKey]*Route
	globalPreMiddleware  []MiddlewareFunc
	globalPostMiddleware []MiddlewareFunc
}

func NewRouter() *Router {
	router := &Router{
		routes: make(map[routeKey]*Route),
		globalPreMiddleware: make([]MiddlewareFunc, 0),
		globalPostMiddleware: make([]MiddlewareFunc, 0),
	}
	return router
}

func (r *Router) MergeRouter(other *Router) {
    for key, route := range other.routes {
        r.routes[key] = route.UsePreMiddlewares(other.globalPreMiddleware).UsePostMiddlewares(other.globalPostMiddleware)
    }
}

func (r *Router) addRoute(method Method, path string, handler func(req *Request, res *Response)) *Route {
	route := routeKey{Method: method, Path: path}
	r.routes[route] = &Route{
		handler:        handler,
		preMiddleware:  make([]MiddlewareFunc, 0),
		postMiddleware: make([]MiddlewareFunc, 0),
	}
	return r.routes[route]
}

func (r *Router) Get(path string, handler func(req *Request, res *Response)) *Route{
	return r.addRoute(GET, path, handler)
}

func (r *Router) Post(path string, handler func(req *Request, res *Response)) *Route{
	return r.addRoute(POST, path, handler)
}

func (r *Router) Put(path string, handler func(req *Request, res *Response)) *Route{
	return r.addRoute(PUT, path, handler)
}

func (r *Router) Patch(path string, handler func(req *Request, res *Response)) *Route{
	return r.addRoute(PATCH, path, handler)
}

func (r *Router) Delete(path string, handler func(req *Request, res *Response)) *Route{
	return r.addRoute(DELETE, path, handler)
}

// UseGlobalPreMiddlewares adds one or more middlewares to run before the handler does run for all routes registered in a specific router
func (r *Router) UseGlobalPreMiddlewares(middlewares []MiddlewareFunc) *Router{
	r.globalPreMiddleware = append(r.globalPreMiddleware, middlewares...)
	return r
}

// UseGlobalPostMiddlewares adds one or more middlewares to run after the handler does run for all routes registered in a specific router
func (r *Router) UseGlobalPostMiddlewares(middlewares []MiddlewareFunc) *Router{
	r.globalPostMiddleware = append(r.globalPostMiddleware, middlewares...)
	return r
}

func (r *Router) Resolve(req *Request, res *Response) {
	routeKey := routeKey{Method: req.GetMethod(), Path: req.GetPath()}

	route, exists := r.routes[routeKey]

	// Execute global-pre-middlewares
	for _, middleware := range r.globalPreMiddleware {
		hasNextBeenCalled := false
		middleware(req, res, func() {
			hasNextBeenCalled = true
		})
		if !hasNextBeenCalled {
			return // Middleware chain was broken
		}
	}

	if !exists {
		res.NotFound()
	} else {
		// Execute pre-middleware for the specific route
		for _, middleware := range route.preMiddleware {
			hasNextBeenCalled := false
			middleware(req, res, func() {
				hasNextBeenCalled = true
			})
			if !hasNextBeenCalled {
				return // Middleware chain was broken
			}
		}

		// Execute the handler
		route.handler(req, res)

		// Execute post-middleware for the specific route
		for _, middleware := range route.postMiddleware {
			hasNextBeenCalled := false
			middleware(req, res, func() {
				hasNextBeenCalled = true
			})
			if !hasNextBeenCalled {
				return // Middleware chain was broken
			}
		}
	}

	// Execute global-post-middleware 
	for _, middleware := range r.globalPostMiddleware {
		hasNextBeenCalled := false
		middleware(req, res, func() {
			hasNextBeenCalled = true
		})
		if !hasNextBeenCalled {
			return // Middleware chain was broken
		}
	}
}
