package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://xvideos.com",
	}
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoPaso := time.Since(inicio)
	//hasta aca tomo el tiempo que tardo en revisar cada pagina, casa uno puede responder en tiempos diferentes y en el for
	//se queda esperando en orden hasta que responda un servidor.... esto es cortante, es decir te corta el flujo si un servidor tarda demasiado
	
	fmt.Printf("Tiempo de ejecucion: %s \n", tiempoPaso)
}

func revisarServidor(servidor string) {
	_, error := http.Get(servidor)
	if error != nil {
		fmt.Println(servidor, "no esta disponible :/")
	} else {
		fmt.Println(servidor, "funca bien")
	}
}

//Goroutines para solucionar que Get se queda ejecutando algun servidor, quiere hacer recurrente(~paral)
