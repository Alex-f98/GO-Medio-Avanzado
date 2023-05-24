package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	respuesta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		panic("Error al leer google")
	}
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)
	/*
		Se utiliza la función io.Copy para copiar los datos de respuesta.Body (el cuerpo de la respuesta HTTP)
		en la instancia e de escritorWeb. io.Copy lee de respuesta.Body y llama al método Write de e
		repetidamente hasta que no hay más datos que copiar.
	*/

}
