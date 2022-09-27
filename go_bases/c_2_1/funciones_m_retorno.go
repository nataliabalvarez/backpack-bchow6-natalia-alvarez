package main

import (
	"errors"
	"fmt"
)

// func calcular(valor1, valor2 float64) (float64, float64, float64, float64) {
// 	suma := valor1 + valor2
// 	resta := valor1 - valor2
// 	multiplicacion := valor1 * valor2
// //	if valor2 != 0 {
// 		divis := valor1 / valor2

// 	return suma, resta, multiplicacion, divis

// }

func calcular(valor1, valor2 float64) (suma, resta, multiplicacion, divis float64) {
	suma = valor1 + valor2
	resta = valor1 - valor2
	multiplicacion = valor1 * valor2
	if valor2 != 0 {
		divis = valor1 / valor2
	}

	return // el return solo, garantiza el orden de las variables

}

func dividir(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("El divisor no puede ser cero")
	}
	return dividendo / divisor, nil
}

func main() {
	// fmt.Println(
	// 	calcular(10, 2),
	// )
	// result, err := dividir(2, 1)
	// if err != nil {
	// 	//tratar el error
	// } else {
	// 	fmt.Println(result)
	// }

	suma, resta, multiplicacion, division := calcular(10, 2)

	fmt.Println("La suma es: ", suma)
	fmt.Println("La resta es: ", resta)
	fmt.Println("La multiplicacion es: ", multiplicacion)
	fmt.Println("La division es: ", division)

}
