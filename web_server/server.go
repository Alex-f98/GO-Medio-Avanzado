package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	srv := new(Server)
	srv.port = port
	srv.router = NewRouter()
	return srv
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	return http.ListenAndServe(s.port, nil)
}

//Proximo curso: https://platzi.com/clases/1846-programacion-golang-2020/26773-que-es-un-middleware/
