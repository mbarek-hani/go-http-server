package main

import (
	"http-server/app"
	"http-server/app/http"
	"http-server/app/middleware"
)

func main() {

	app := app.Application()

	app.Get("/", handleHome)
	app.Get("/about", handleAbout)
	app.Get("/contact", handleContact)

	app.AddRouterPostMiddleware(middleware.LoggerMiddleware)

	app.Start("8000")

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleHome(req *http.Request, res *http.Response) {
	res.JsonResponse(Person{"Jhon Doe", 18})
}

func handleAbout(req *http.Request, res *http.Response) {
	res.HttpResponse("Welcome to about")
}

func handleContact(req *http.Request, res *http.Response) {
	res.HttpResponse("Welcome to contact")
}
