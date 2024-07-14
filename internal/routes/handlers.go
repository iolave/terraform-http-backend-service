package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/iolave/terraform-http-backend-service/internal/errors"
	"github.com/iolave/terraform-http-backend-service/internal/storage"
)

func getTfStateHandler(w http.ResponseWriter, r *http.Request) {
	storagePath, httpErr := getStoragePath()

	if httpErr != nil {
		w.WriteHeader(httpErr.StatusCode)
		w.Write(httpErr.ToJsonBytes())
		return
	}

	appId := chi.URLParam(r, "projectId")
	env := chi.URLParam(r, "env")

	storage := storage.NewLocalStorage(storagePath)

	content, err := storage.ReadTfState(appId, env)

	if err != nil {
		httpErr := errors.HttpInternalServerError(err.Error())
		w.WriteHeader(httpErr.StatusCode)
		w.Write(httpErr.ToJsonBytes())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func postTfStateHandler(w http.ResponseWriter, r *http.Request) {
	storagePath, httpErr := getStoragePath()

	if httpErr != nil {
		w.WriteHeader(httpErr.StatusCode)
		w.Write(httpErr.ToJsonBytes())
		return
	}

	appId := chi.URLParam(r, "projectId")
	env := chi.URLParam(r, "env")

	if r.Body == nil {
		httpErr := errors.HttpBadRequestError("no request body found")
		w.WriteHeader(httpErr.StatusCode)
		w.Write(httpErr.ToJsonBytes())
		return
	}

	content, _ := io.ReadAll(r.Body)
	r.Body.Close()

	storage := storage.NewLocalStorage(storagePath)

	err := storage.WriteTfState(appId, env, content)

	if err != nil {
		httpErr := errors.HttpInternalServerError(err.Error())
		w.WriteHeader(httpErr.StatusCode)
		w.Write(httpErr.ToJsonBytes())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func getStoragePath() (string, *errors.HttpError) {
	path, present := os.LookupEnv("STORAGE_PATH")

	if !present {
		return "", errors.HttpInternalServerError("STORAGE_PATH environment varaible is not set")
	}

	return path, nil
}
