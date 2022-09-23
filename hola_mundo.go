package main

import (
	"fmt"
)

var tiempo int //inicializado
const segundos int = 10

func main() {
	fmt.Println("Hola mundis")
	horas := 25.0           // inicializado y asignacion
	fmt.Printf("%T", horas) //imnprime el tipo de variable

	tiempo = 2 // asignacion
	fmt.Println("\n", tiempo)

	fmt.Println(segundos)
}
