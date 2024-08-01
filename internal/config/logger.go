package config

import "github.com/iolave/go-logger"

var Logger = logger.New("terraform-http-backend-service")

type Trace struct {
	RequestId string `json:"requestId"`
}
