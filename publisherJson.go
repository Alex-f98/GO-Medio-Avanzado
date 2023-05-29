package main3

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	// URL de destino y cuerpo de la solicitud
	url := "http://localhost:8080"
	contentType := "/"
	jsonStr := []byte(`{"script": "my_script.py"}`)

	// Realizar la solicitud HTTP POST
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}
	defer resp.Body.Close()
}

/*
 Es importante tener en cuenta que http.Post utiliza una instancia de http.DefaultClient para enviar la solicitud.
 Si necesitas personalizar el cliente HTTP, puedes crear un cliente personalizado utilizando http.Client y
 utilizar su método Post.
*/
