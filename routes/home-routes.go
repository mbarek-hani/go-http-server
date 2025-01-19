package routes

import (
	"http-server/app/http"
	"http-server/controllers/HomeController"
)

func HomeRoutes() *http.Router {
	router := http.NewRouter()

	router.Get("/home", HomeController.Index)

	return router
}