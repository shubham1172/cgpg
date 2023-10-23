package generated

type HTTP2Codes struct {
	Code    string
	Message string
}

func (e *HTTP2Codes) GetMessage() string {
	return e.Message
}

func (e *HTTP2Codes) GetCode() string {
	return e.Code
}

var (
	HTTP2_200 = &HTTP2Codes{
		Code:    "HTTP2_200",
		Message: "Status code 2 200 OK",
	}

	HTTP2_201 = &HTTP2Codes{
		Code:    "HTTP2_201",
		Message: "Status code 2 201 Created",
	}

	HTTP2_202 = &HTTP2Codes{
		Code:    "HTTP2_202",
		Message: "Status code 2 202 Accepted",
	}

	HTTP2_301 = &HTTP2Codes{
		Code:    "HTTP2_301",
		Message: "Status code 2 301 Moved Permanently",
	}

	HTTP2_302 = &HTTP2Codes{
		Code:    "HTTP2_302",
		Message: "Status code 2 302 Found",
	}

	HTTP2_400 = &HTTP2Codes{
		Code:    "HTTP2_400",
		Message: "Status code 2 400 Bad Request",
	}

	HTTP2_401 = &HTTP2Codes{
		Code:    "HTTP2_401",
		Message: "Status code 2 401 Unauthorized",
	}

	HTTP2_403 = &HTTP2Codes{
		Code:    "HTTP2_403",
		Message: "Status code 2 403 Forbidden",
	}

	HTTP2_404 = &HTTP2Codes{
		Code:    "HTTP2_404",
		Message: "Status code 2 404 Not Found",
	}

	HTTP2_500 = &HTTP2Codes{
		Code:    "HTTP2_500",
		Message: "Status code 2 500 Internal Server Error",
	}
)
