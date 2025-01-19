package main

import (
	"http-server/app"
	"http-server/controllers/AboutController"
	"http-server/controllers/ContactController"
	"http-server/routes"
)

func main() {

	app := app.Application()

	app.Add(routes.HomeRoutes())

	app.Get("/about", AboutController.Index)
	app.Get("/contact", ContactController.Index)

	app.Start("8000")

}
