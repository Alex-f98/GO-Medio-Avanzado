package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Código a ejecutar en Python
	code := `
def add(a, b):
    return a + b

result = add(5, 3)
print(result)
`

	// URL del servicio web de Programiz
	url := "https://www.programiz.com/api/compile-python"

	// Preparar el cuerpo de la solicitud
	body := bytes.NewBufferString(code)

	// Realizar la solicitud HTTP POST
	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Imprimir el resultado
	fmt.Println("Respuesta:", string(responseBody))
}
