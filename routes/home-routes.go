package routes

import (
	"http-server/app/http"
	"http-server/controllers/HomeController"
	"http-server/middleware"
)

func HomeRoutes() *http.Router {
	router := http.NewRouter()

	router.Get("/home", HomeController.Index)
	router.Get("/home/1", HomeController.Index)
	router.Get("/home/2", HomeController.Index)
	router.Get("/home/3", HomeController.Index)

	router.UseGlobalPreMiddlewares([]http.MiddlewareFunc{middleware.AuthMiddleware})
	return router
}