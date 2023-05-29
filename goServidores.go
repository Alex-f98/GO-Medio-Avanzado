package main2

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://xvideos.com",
	}

	for {
		for _, servidor := range servidores {
			go _revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
	}
	time.Since(inicio)
}
func _revisarServidor(servidor string, canal chan string) {
	_, error := http.Get(servidor)
	if error != nil {
		canal <- servidor + " no esta disponible :o" //canal toma lo que esta a derecha y transmite
	} else {
		canal <- servidor + " todo funcina bien :D"
	}
}

/*













 */

/*
func main() {
	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://xvideos.com",
	}
	for _, servidor := range servidores {
		//go va a crear un hilo de ejecucion por cada uno de estos servidorres independientemente de main()
		go _revisarServidor(servidor, canal)
		//el main() simplemente crea las 4 subrutinas y termina, luego las subrutinas se pierden
	}
	//fmt.Println(<-canal)
	//en principio queda esperando al primero que responda y printea, google es mas rapido.
	//http://google.com funca bien
	//http://google.comtodo funcina bien :D

	//listo, ahora
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion: %s \n", tiempoPaso) //Tiempo de ejecucion: 0s (x q se ejecuta por fuera del main())
}

func _revisarServidor(servidor string, canal chan string) {
	_, error := http.Get(servidor)
	if error != nil {
		fmt.Println(servidor, "no esta disponible :/")
		canal <- servidor + " no esta disponible :o" //canal toma lo que esta a derecha y transmite
	} else {
		fmt.Println(servidor, "funca bien")
		canal <- servidor + " todo funcina bien :D"
	}
}
*/
//Goroutines para solucionar que Get se queda ejecutando algun servidor, quiere hacer recurrente(~paral)
