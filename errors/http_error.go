package errors

// NewHTTPError returns http status code and error message
func NewHTTPError(statusCode int, message string) error {
	return &httpError{
		statusCode: statusCode,
		message:    message,
	}
}

type httpError struct {
	statusCode int
	message    string
}

func (e *httpError) Error() string {
	return e.message
}

func (e *httpError) GetStatusCode() int {
	return e.statusCode
}

func (e *httpError) GetMessage() string {
	return e.message
}

// HTTPError describes the interface for http error
type HTTPError interface {
	GetStatusCode() int
	GetMessage() string
}
