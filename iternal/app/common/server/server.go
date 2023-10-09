package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Run(host string, port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: handler,
	}

	return server.httpServer.ListenAndServe()
}
