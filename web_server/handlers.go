package main

import (
	"fmt"
	"net/http"
)

func HnadLeRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "QUE MIRA BOBO X2!!")
}
