package http

import "strings"

type Method int

const (
	InvalidMethod Method = iota
	GET
	POST
	PUT
	DELETE
	PATCH
)

func (m Method) String() string {
	return []string{"InvalidMethod", "GET", "POST", "PUT", "DELETE", "PATCH"}[m]
}

func ParseToMethod(method string) Method {
	switch strings.ToUpper(method) {
	case "GET":
		return GET
	case "POST":
		return POST
	case "PUT":
		return PUT
	case "DELETE":
		return DELETE
	case "PATCH":
		return PATCH
	default:
		return InvalidMethod
	}
}
