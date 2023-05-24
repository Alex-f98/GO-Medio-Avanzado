package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	rtr := new(Router)
	rtr.rules = make(map[string]http.HandlerFunc, 0)
	return rtr
}

func (rtr *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	//si un servidor no ´puede devolver la ruta, deberia retornar un 404.
	handler, exist := rtr.rules[path]
	return handler, exist
}

// HASTA ACA NOS QUEDAMOS WEY!: https://platzi.com/clases/1846-programacion-golang-2020/26774-manejando-request-http/
func (rtr *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "QUE MIRA BOBO")
}
