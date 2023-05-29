package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Params struct {
	Scripts []Script `json:"scripts"`
}

type Script struct {
	Lang string `json:"lang"`
	Path string `json:"path"`
}

func main() {
	pattern := "/"
	http.HandleFunc(pattern, handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var params Params
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &params)
	if err != nil {
		fmt.Println(err)
		return
	}

	for idx, script := range params.Scripts {
		fmt.Println(idx)
		go execScript(script)
	}
}

func execScript(script Script) {
	var cmd *exec.Cmd
	/*
		fmt.Println("_script: ", script)
		fmt.Println("_lang: ", script.Lang)
		fmt.Println("_path: ", script.Path)
		fmt.Println("\n")
	*/

	switch script.Lang {
	//
	case "python2":
		fmt.Println("Ejecutando PYTHON2")
		cmd = exec.Command("python", script.Path)
		//
	case "python3":
		fmt.Println("Ejecutando PYTHON3")
		cmd = exec.Command("python", script.Path)
		//
	default:
		fmt.Println("Este lenguage no es de programacion, marmota")
		//
		return
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out.String())
}

/*
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Params struct {
	Scripts []Script `json:"scripts"`
}

type Script struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func main() {
	// Se utiliza para registrar un manejador de solicitudes HTTP para un determinado patrón de ruta.
	pattern := "/"
	http.HandleFunc(pattern, handler)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", nil)
}

// handler es una función que se ejecutará cuando se reciba una solicitud que coincida con el patrón especificado.
// La función debe tener la siguiente firma: func(ResponseWriter, *Request), donde
// ResponseWriter se utiliza para enviar la respuesta HTTP al cliente y
// *Request contiene la información de la solicitud recibida.
func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var params Params
	// El json esta en body y se guarada en la estructura Params.
	err = json.Unmarshal(body, &params)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, script := range params.Script {
		go execCodes(script)
	}

}

func execCodes(script Script) {
	var out bytes.Buffer
	var cmd *exec.Cmd
	cmd.Stdout = &out

	switch script.Code {
	case "Python1":
		cmd = exec.Command("python", script.Script)
	case "Python2":
		cmd = exec.Command("python", script.Script)
	default:
		fmt.Println("Che! que te equivocaste!")
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func _execCodes(params *Params) {
	//Ejecuta usando : "python <params.Script>"
	cmd := exec.Command("python", params.Script)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out.String())
}
*/
/*





package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//cmd := exec.Command("Rscript", "myscript.R")
	cmd := exec.Command("python", "my_script.py")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out.String())
}
*/
