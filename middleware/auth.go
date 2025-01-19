package middleware

import (
	"http-server/app/http"
	"log"
)

func AuthMiddleware(req *http.Request, res *http.Response, next func()) {
	if req.GetHeader("Authorisation") == "" {
		res.HttpResponse("Unauthorized", http.StatusUnauthorized)
		log.Print(req.GetMethod(), " ", req.GetPath(), " ", res.GetStatusCode().Int(), " ", res.GetStatusCode())
		return
	}
	next()
}