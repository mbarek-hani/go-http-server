package http

// StatusCode represents HTTP status codes
type StatusCode int

const (
	// 1xx Informational
	StatusContinue           StatusCode = 100
	StatusSwitchingProtocols StatusCode = 101
	StatusProcessing         StatusCode = 102
	StatusEarlyHints         StatusCode = 103

	// 2xx Success
	StatusOK                   StatusCode = 200
	StatusCreated              StatusCode = 201
	StatusAccepted             StatusCode = 202
	StatusNonAuthoritativeInfo StatusCode = 203
	StatusNoContent            StatusCode = 204
	StatusResetContent         StatusCode = 205
	StatusPartialContent       StatusCode = 206
	StatusMultiStatus          StatusCode = 207
	StatusAlreadyReported      StatusCode = 208
	StatusIMUsed               StatusCode = 226

	// 3xx Redirection
	StatusMultipleChoices   StatusCode = 300
	StatusMovedPermanently  StatusCode = 301
	StatusFound             StatusCode = 302
	StatusSeeOther          StatusCode = 303
	StatusNotModified       StatusCode = 304
	StatusUseProxy          StatusCode = 305
	StatusTemporaryRedirect StatusCode = 307
	StatusPermanentRedirect StatusCode = 308

	// 4xx Client Error
	StatusBadRequest                   StatusCode = 400
	StatusUnauthorized                 StatusCode = 401
	StatusPaymentRequired              StatusCode = 402
	StatusForbidden                    StatusCode = 403
	StatusNotFound                     StatusCode = 404
	StatusMethodNotAllowed             StatusCode = 405
	StatusNotAcceptable                StatusCode = 406
	StatusProxyAuthRequired            StatusCode = 407
	StatusRequestTimeout               StatusCode = 408
	StatusConflict                     StatusCode = 409
	StatusGone                         StatusCode = 410
	StatusLengthRequired               StatusCode = 411
	StatusPreconditionFailed           StatusCode = 412
	StatusRequestEntityTooLarge        StatusCode = 413
	StatusRequestURITooLong            StatusCode = 414
	StatusUnsupportedMediaType         StatusCode = 415
	StatusRequestedRangeNotSatisfiable StatusCode = 416
	StatusExpectationFailed            StatusCode = 417
	StatusTeapot                       StatusCode = 418
	StatusMisdirectedRequest           StatusCode = 421
	StatusUnprocessableEntity          StatusCode = 422
	StatusLocked                       StatusCode = 423
	StatusFailedDependency             StatusCode = 424
	StatusTooEarly                     StatusCode = 425
	StatusUpgradeRequired              StatusCode = 426
	StatusPreconditionRequired         StatusCode = 428
	StatusTooManyRequests              StatusCode = 429
	StatusRequestHeaderFieldsTooLarge  StatusCode = 431
	StatusUnavailableForLegalReasons   StatusCode = 451

	// 5xx Server Error
	StatusInternalServerError           StatusCode = 500
	StatusNotImplemented                StatusCode = 501
	StatusBadGateway                    StatusCode = 502
	StatusServiceUnavailable            StatusCode = 503
	StatusGatewayTimeout                StatusCode = 504
	StatusHTTPVersionNotSupported       StatusCode = 505
	StatusVariantAlsoNegotiates         StatusCode = 506
	StatusInsufficientStorage           StatusCode = 507
	StatusLoopDetected                  StatusCode = 508
	StatusNotExtended                   StatusCode = 510
	StatusNetworkAuthenticationRequired StatusCode = 511
)

// String returns the string representation of the status code
func (c StatusCode) String() string {
	switch c {
	case StatusContinue:
		return "Continue"
	case StatusSwitchingProtocols:
		return "Switching Protocols"
	case StatusProcessing:
		return "Processing"
	case StatusEarlyHints:
		return "Early Hints"
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusAccepted:
		return "Accepted"
	case StatusNonAuthoritativeInfo:
		return "Non-Authoritative Information"
	case StatusNoContent:
		return "No Content"
	case StatusResetContent:
		return "Reset Content"
	case StatusPartialContent:
		return "Partial Content"
	case StatusMultiStatus:
		return "Multi-Status"
	case StatusAlreadyReported:
		return "Already Reported"
	case StatusIMUsed:
		return "IM Used"
	case StatusMultipleChoices:
		return "Multiple Choices"
	case StatusMovedPermanently:
		return "Moved Permanently"
	case StatusFound:
		return "Found"
	case StatusSeeOther:
		return "See Other"
	case StatusNotModified:
		return "Not Modified"
	case StatusUseProxy:
		return "Use Proxy"
	case StatusTemporaryRedirect:
		return "Temporary Redirect"
	case StatusPermanentRedirect:
		return "Permanent Redirect"
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusProxyAuthRequired:
		return "Proxy Authentication Required"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "Length Required"
	case StatusPreconditionFailed:
		return "Precondition Failed"
	case StatusRequestEntityTooLarge:
		return "Request Entity Too Large"
	case StatusRequestURITooLong:
		return "Request URI Too Long"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusRequestedRangeNotSatisfiable:
		return "Requested Range Not Satisfiable"
	case StatusExpectationFailed:
		return "Expectation Failed"
	case StatusTeapot:
		return "I'm a teapot"
	case StatusMisdirectedRequest:
		return "Misdirected Request"
	case StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "Failed Dependency"
	case StatusTooEarly:
		return "Too Early"
	case StatusUpgradeRequired:
		return "Upgrade Required"
	case StatusPreconditionRequired:
		return "Precondition Required"
	case StatusTooManyRequests:
		return "Too Many Requests"
	case StatusRequestHeaderFieldsTooLarge:
		return "Request Header Fields Too Large"
	case StatusUnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusGatewayTimeout:
		return "Gateway Timeout"
	case StatusHTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	case StatusVariantAlsoNegotiates:
		return "Variant Also Negotiates"
	case StatusInsufficientStorage:
		return "Insufficient Storage"
	case StatusLoopDetected:
		return "Loop Detected"
	case StatusNotExtended:
		return "Not Extended"
	case StatusNetworkAuthenticationRequired:
		return "Network Authentication Required"
	default:
		return "Unknown Status Code"
	}
}

// Int returns the integer value of the status code
func (c StatusCode) Int() int {
	return int(c)
}

// IsInformational returns true if status code is 1xx
func (c StatusCode) IsInformational() bool {
	return c >= 100 && c < 200
}

// IsSuccess returns true if status code is 2xx
func (c StatusCode) IsSuccess() bool {
	return c >= 200 && c < 300
}

// IsRedirection returns true if status code is 3xx
func (c StatusCode) IsRedirection() bool {
	return c >= 300 && c < 400
}

// IsClientError returns true if status code is 4xx
func (c StatusCode) IsClientError() bool {
	return c >= 400 && c < 500
}

// IsServerError returns true if status code is 5xx
func (c StatusCode) IsServerError() bool {
	return c >= 500 && c < 600
}

// IsError returns true if status code is 4xx or 5xx
func (c StatusCode) IsError() bool {
	return c >= 400 && c < 600
}
