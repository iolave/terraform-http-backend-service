package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	TF_STATE_PATH = "/tf-state/{projectId}/{env}"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.SetHeader("content-type", "application/json"))

	router.Get(TF_STATE_PATH, getTfStateHandler)
	router.Post(TF_STATE_PATH, postTfStateHandler)

	return router
}
