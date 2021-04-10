package controllers

import (
	"github.com/gorilla/pat"
	"net/http"
)

type Server struct {
	Handler *pat.Router
}

func (s *Server) HandleGetRequestPath(path string, handlerFunc http.HandlerFunc) {
	s.Handler.Get(path, handlerFunc)
}

func (s *Server) HandleDeleteRequestPath(path string, handlerFunc http.HandlerFunc) {
	s.Handler.Delete(path, handlerFunc)
}

func (s *Server) HandlePutRequestPath(path string, handlerFunc http.HandlerFunc) {
	s.Handler.Put(path, handlerFunc)
}

func (s *Server) HandlePostRequestPath(path string, handlerFunc http.HandlerFunc) {
	s.Handler.Post(path, handlerFunc)
}
