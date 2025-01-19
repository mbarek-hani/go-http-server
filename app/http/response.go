package http

import (
	"fmt"
	"http-server/helpers"
	"strconv"
	"time"
)

type Response struct {
	statusCode StatusCode
	headers    map[string]string
	body       string
}

func NewHttpResponse() *Response {
	response := &Response{}
	response.headers = make(map[string]string)
	return response
}

func (r *Response) SetHeader(headerName string, headerValue string) {
	r.headers[headerName] = headerValue
}

func (r *Response) SetStatusCode(code StatusCode) {
	r.statusCode = code
}

func (r *Response) GetStatusCode() StatusCode {
	return r.statusCode
}

func (r *Response) Redirect(url string) {
	r.SetStatusCode(StatusSeeOther)
	r.SetHeader("Location", url)
}

func (r *Response) String() string {
	var rawResponse string
	rawResponse = fmt.Sprintf("HTTP/1.1 %v %v\r\n", r.statusCode.Int(), r.statusCode)
	for headerName, headerValue := range r.headers {
		rawResponse += fmt.Sprintf("%v: %v\r\n", headerName, headerValue)
	}
	rawResponse += fmt.Sprintf("\r\n%v", r.body)
	return rawResponse
}

func (r *Response) JsonResponse(payload interface{}) {
	contentLength, body := helpers.MustToJSONString(payload)
	r.SetStatusCode(StatusOK)
	r.SetHeader("Date", time.Now().UTC().Format(time.RFC1123))
	r.SetHeader("Server", "GoHTTP/1.0")
	r.SetHeader("Connection", "close")
	r.SetHeader("Content-Type", "application/json; charset=utf-8")
	r.SetHeader("Content-Length", strconv.Itoa(contentLength))
	r.body = body
}

func (r *Response) HttpResponse(payload string) {
	r.SetStatusCode(StatusOK)
	r.SetHeader("Date", time.Now().UTC().Format(time.RFC1123))
	r.SetHeader("Server", "GoHTTP/1.0")
	r.SetHeader("Connection", "close")
	r.SetHeader("Content-Type", "text/plain; charset=utf-8")
	r.SetHeader("Content-Length", strconv.Itoa(len(payload)))
	r.body = payload
}

func (r *Response) NotFound() {
	r.SetStatusCode(StatusNotFound)
	r.SetHeader("Date", time.Now().UTC().Format(time.RFC1123))
	r.SetHeader("Server", "GoHTTP/1.0")
	r.SetHeader("Connection", "close")
	r.body = ""
}
