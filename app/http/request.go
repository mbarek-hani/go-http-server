package http

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Request struct {
	method      Method
	path        string
	queryParams map[string]string
	headers     map[string]string
	body        map[string]string
}

func ParseToRequest(rawRequest []byte) (*Request, error) {
	request := &Request{}

	// Check for empty request
	if len(rawRequest) == 0 {
		return nil, fmt.Errorf("empty request")
	}

	parts := strings.Split(string(rawRequest), "\r\n\r\n")
	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid request format")
	}

	headerSection := parts[0]
	lines := strings.Split(headerSection, "\r\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("missing request line")
	}

	// Parse first line (request line)
	firstLine := strings.Split(lines[0], " ")
	if len(firstLine) < 3 {
		return nil, fmt.Errorf("invalid request line format")
	}

	// Parse method
	method := ParseToMethod(firstLine[0])
	if method == InvalidMethod {
		return nil, fmt.Errorf("invalid HTTP method: %s", firstLine[0])
	}
	request.method = method

	// Parse path and query
	pathAndQuery := strings.Split(firstLine[1], "?")
	if len(pathAndQuery[0]) == 0 {
		request.path = "/"
	} else {
		request.path = pathAndQuery[0]
	}

	// Parse query parameters
	if len(pathAndQuery) > 1 {
		request.queryParams = make(map[string]string)
		queries := strings.Split(pathAndQuery[1], "&")
		for _, pair := range queries {
			keyValue := strings.Split(pair, "=")
			if len(keyValue) == 2 {
				request.queryParams[keyValue[0]] = keyValue[1]
			}
		}
	}

	// Parse headers
	request.headers = make(map[string]string)
	if len(lines) > 1 {
		for _, header := range lines[1:] {
			if header == "" {
				continue
			}
			keyValue := strings.Split(header, ": ")
			if len(keyValue) != 2 {
				return nil, fmt.Errorf("invalid header format: %s", header)
			}
			request.headers[keyValue[0]] = keyValue[1]
		}
	}

	// Parse body
	if len(parts) > 1 {
		body, err := request.parseBody([]byte(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("parsing body: %w", err)
		}
		request.body = body
	}

	return request, nil
}

// func ParseToRequest(rawRequest []byte) *HttpRequest {
//     request := &HttpRequest{}

//     parts := strings.Split(string(rawRequest), "\r\n\r\n")
//     headerSection := parts[0]

//     lines := strings.Split(headerSection, "\r\n")
//     firstLine := strings.Split(lines[0], " ")
//     request.method = ParseToMethod(firstLine[0])

//     pathAndQuery := strings.Split(firstLine[1], "?")
//     request.path = pathAndQuery[0]

//     if len(pathAndQuery) > 1 {
//         queries := strings.Split(pathAndQuery[1], "&")
//         request.queryParams = make(map[string]string)
//         for _, pair := range queries {
//             keyValue := strings.Split(pair, "=")
//             if len(keyValue) == 2 {
//                 request.queryParams[keyValue[0]] = keyValue[1]
//             }
//         }
//     }

//     request.headers = make(map[string]string)
//     if len(lines) > 1 {
//         for _, header := range lines[1:] {
//             if header == "" {
//                 continue
//             }
//             keyValue := strings.Split(header, ": ")
//             if len(keyValue) == 2 {
//                 request.headers[keyValue[0]] = keyValue[1]
//             }
//         }
//     }

// 	if len(parts) > 1 {
//         request.body = request.parseBody([]byte(parts[1]))
//     }

//     if request.path == "" {
//         request.path = "/"
//     }

//     return request
// }

func (r *Request) GetMethod() Method {
	return r.method
}

func (r *Request) GetPath() string {
	return r.path
}

func (r *Request) GetQueryParam(key string) string {
	return r.queryParams[key]
}

func (r *Request) GetHeader(headerName string) string {
	return r.headers[headerName]
}

func (r *Request) parseBody(body []byte) (map[string]string, error) {
	contentType := r.headers["Content-Type"]

	switch contentType {
	case "application/json":
		var result map[string]string
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return result, nil

	case "application/x-www-form-urlencoded":
		formData := make(map[string]string)
		pairs := strings.Split(string(body), "&")
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				formData[kv[0]] = kv[1]
			}
		}
		return formData, nil

	default:
		return r.body, nil
	}
}
