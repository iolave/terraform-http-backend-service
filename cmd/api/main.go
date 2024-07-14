package main

import "github.com/iolave/terraform-http-backend-service/internal/server"

func main() {
	srv := server.NewServer("0.0.0.0", 3000)

	srv.Serve()
}
