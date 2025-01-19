package HomeController

import "http-server/app/http"

func Index(req *http.Request, res *http.Response) {
		res.HttpResponse("Hello from home page")
}