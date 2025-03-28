package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port            string
	middlewareFuncs []func(http.Handler) http.Handler
	mux             *http.ServeMux
	handlers        []Handler
}

type Handler struct {
	Path        string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

// Instanciates a new Server that is essentially a wrapper for a mux server.
func NewServer(mux *http.ServeMux, port string) Server {
	return Server{
		port: port,
		mux:  mux,
	}
}

// Appends a new middleware function to the server
func (s *Server) WithMiddlewareFunc(middlewareFunc func(http.Handler) http.Handler) {
	s.middlewareFuncs = append(s.middlewareFuncs, middlewareFunc)
}

// Appends a new handler function to the server
func (s *Server) WithHandler(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.handlers = append(s.handlers, Handler{path, handler})
}

// PRIVATE
// executes each middleware function then returns the finished handler for the handler to run.
func (s *Server) executeMiddleware(h http.Handler) http.Handler {
	for _, middlewareFunc := range s.middlewareFuncs {
		h = middlewareFunc(h)
	}

	return h
}

// Builds the mux server needed to serve.
// Iterates through each handler and hooks the chain of middleware functions to the handler.
func (s *Server) Serve() {
	for _, handler := range s.handlers {
		s.mux.Handle(handler.Path, s.executeMiddleware(http.HandlerFunc(handler.HandlerFunc)))
	}

	fmt.Printf("Starting server at port %s\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.mux))
}
