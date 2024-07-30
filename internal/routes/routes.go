package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iolave/terraform-http-backend-service/internal/config"
)

const (
	TF_STATE_PATH = "/tf-state/{projectId}/{env}"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(requestLoggerMdw(config.Logger))
	router.Use(middleware.SetHeader("content-type", "application/json"))

	router.Get(TF_STATE_PATH, getTfStateHandler)
	router.Post(TF_STATE_PATH, postTfStateHandler)

	return router
}
