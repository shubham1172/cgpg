
package generated

type HTTPCodes struct {
	Code string
	Message string
}

func (e *HTTPCodes) GetMessage() string {
	return e.Message
}

func (e *HTTPCodes) GetCode() string {
	return e.Code
}

var (
	
	HTTP_200 = &HTTPCodes{
		Code: "HTTP_200",
		Message: "Status code 200 OK",
	}
	
	HTTP_201 = &HTTPCodes{
		Code: "HTTP_201",
		Message: "Status code 201 Created",
	}
	
	HTTP_202 = &HTTPCodes{
		Code: "HTTP_202",
		Message: "Status code 202 Accepted",
	}
	
	HTTP_301 = &HTTPCodes{
		Code: "HTTP_301",
		Message: "Status code 301 Moved Permanently",
	}
	
	HTTP_302 = &HTTPCodes{
		Code: "HTTP_302",
		Message: "Status code 302 Found",
	}
	
	HTTP_400 = &HTTPCodes{
		Code: "HTTP_400",
		Message: "Status code 400 Bad Request",
	}
	
	HTTP_401 = &HTTPCodes{
		Code: "HTTP_401",
		Message: "Status code 401 Unauthorized",
	}
	
	HTTP_403 = &HTTPCodes{
		Code: "HTTP_403",
		Message: "Status code 403 Forbidden",
	}
	
	HTTP_404 = &HTTPCodes{
		Code: "HTTP_404",
		Message: "Status code 404 Not Found",
	}
	
	HTTP_500 = &HTTPCodes{
		Code: "HTTP_500",
		Message: "Status code 500 Internal Server Error",
	}
	
)