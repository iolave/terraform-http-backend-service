package errors

import "net/http"

func HttpNotFoundError(details string) *HttpError {
	return &HttpError{
		Name:       "not_found_error",
		Details:    details,
		StatusCode: http.StatusNotFound,
	}
}

func HttpBadRequestError(details string) *HttpError {
	return &HttpError{
		Name:       "bad_request_error",
		Details:    details,
		StatusCode: http.StatusBadRequest,
	}
}

func HttpInternalServerError(details string) *HttpError {
	return &HttpError{
		Name:       "internal_server_error",
		Details:    details,
		StatusCode: http.StatusInternalServerError,
	}
}
