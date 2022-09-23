package main

import "fmt"

var nombre string = "Natalia"
var direccion string
var (
	temp    = 15.3
	humedad = 56
	presion = 1100
)

func main() {
	fmt.Println("\n--------- Ejercicio 1 ---------")
	fmt.Println("Nombre: ", nombre)

	direccion = "Ricardo Gutierrez"

	fmt.Println("Direccion: ", direccion)

	fmt.Println("\n--------- Ejercicio 2 ---------")
	fmt.Println("Temperatura: ", temp)
	fmt.Printf("%T", temp)
	fmt.Println()
	fmt.Println("Humedad: ", humedad)
	fmt.Printf("%T", humedad)
	fmt.Println()
	fmt.Println("Presion: ", presion)
	fmt.Printf("%T", presion)
	fmt.Println()

	fmt.Println("\n--------- Ejercicio 3 ---------")
	//var 1nombre string -- las variables no deben empezar con valores numericos
	//var nombre string -- correccion
	//var apellido string -- ok
	//var int edad -- var nombre tipo
	//var eddad int -- correccion
	//1apellido := 6 -- las variables no deben empezar con valores numericos
	//var licencia_de_conducir = true -- ok
	//var estatura de la persona int -- el nombre de las variables no pueden tener espacios
	//cantidadDeHijos := 2 -- ok

	fmt.Println("\n--------- Ejercicio 4 ---------")
	/*	var apellido string = "Gomez"
		var edad int = "35"
		boolean := "false";
		var sueldo string = 45857.90
		var nombre string = "Julián"
	*/
	var apellido string = "Gomez"
	var edad int = 35
	variable := false
	var sueldo string = "45857.90"
	var nombre string = "Julián"
	fmt.Println(apellido, edad, variable, sueldo, nombre)

}
