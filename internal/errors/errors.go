package errors

import "encoding/json"

type HttpError struct {
	Name       string `json:"name"`
	Details    string `json:"details"`
	StatusCode int    `json:"statusCode"`
}

func (httpErr HttpError) ToJsonBytes() []byte {
	b, _ := json.Marshal(httpErr)

	return b
}
