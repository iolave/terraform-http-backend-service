package server

import (
	"net/http"

	"github.com/iolave/terraform-http-backend-service/internal/errors"
)

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	httpErr := errors.HttpNotFoundError("Resource not found")

	w.WriteHeader(httpErr.StatusCode)
	w.Write(httpErr.ToJsonBytes())
}
