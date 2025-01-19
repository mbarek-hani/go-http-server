package app

import (
	"http-server/app/config"
	"http-server/app/http"
	"log"
)

type App struct {
	server *http.HttpServer
	*http.Router
}

func Application() *App {
	app := &App{}
	app.Router = http.NewRouter()
	return app
}

func (a *App) Start(port string) {
	a.server = http.NewHttpServer("localhost", port)
	a.UseGlobalPreMiddlewares(config.GlobalPreMiddlewares())
	a.UseGlobalPostMiddlewares(config.GlobalPostMiddlewares())
	log.Println("Server is listening on http://localhost:" + a.server.Port)
	a.server.Listen(a.Router)
}

func (a *App) Add(RouteGroup *http.Router) {
	a.MergeRouter(RouteGroup)
}
