package main

import (
	"fmt"
	"net/http"
)

func holaHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hola\n")
	//ctrl+c para interruptir la ejecucion
}

func main(){
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}