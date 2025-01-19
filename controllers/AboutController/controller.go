package AboutController

import (
	"http-server/app/http"
)

func Index(req *http.Request, res *http.Response) {
		res.HttpResponse("Hello from about page")
}