package config

import (
	"http-server/app/http"
	"http-server/middleware"
)

func GlobalPreMiddlewares() []http.MiddlewareFunc {
	return []http.MiddlewareFunc{}
}

func GlobalPostMiddlewares() []http.MiddlewareFunc {
	return []http.MiddlewareFunc{
		middleware.LoggerMiddleware,
	}
}