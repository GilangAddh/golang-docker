package http

type RequestError struct {
	StatusCode int
	Message    string
	Errors     any
}

func (r *RequestError) Error() string {
	return r.Message
}

func NewRequestError(statusCode int, message string, errors any) error {
	return &RequestError{StatusCode: statusCode, Message: message, Errors: errors}
}
