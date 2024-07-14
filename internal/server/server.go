package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iolave/terraform-http-backend-service/internal/routes"
)

type server struct {
	port   int32
	host   string
	router chi.Router
}

func NewServer(host string, port int32) *server {
	// chi.RegisterMethod("LOCK")
	// chi.RegisterMethod("UNLOCK")
	router := routes.NewRouter()

	router.NotFound(notFoundHandler)

	srv := &server{
		port:   port,
		host:   host,
		router: router,
	}

	return srv
}

func (srv *server) Serve() {
	addr := fmt.Sprintf("%s:%d", srv.host, srv.port)

	if err := http.ListenAndServe(addr, srv.router); err != nil {
		log.Fatal("unable to start server", err.Error())
	}
}
