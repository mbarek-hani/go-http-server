package middleware

import (
	"http-server/app/http"
	"log"
)

func LoggerMiddleware(req *http.Request, res *http.Response, next func()) {
	log.Print(req.GetMethod(), " ", req.GetPath(), " ", res.GetStatusCode().Int(), " ", res.GetStatusCode())
	next()
}
