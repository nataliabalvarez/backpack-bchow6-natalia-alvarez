package main

import "fmt"

func inspeccionarVariable(numero int) {
	if numero > 0 {
		fmt.Println("El numero es positivo")
	} else if numero < 0 {
		fmt.Println("El numero  es negativo")

	} else {
		fmt.Println("El numero  es cero")

	}

}

func suma(num1, num2 float64) float64 {
	result := num1 + num2
	return result
}

const (
	OperadorSuma           = "+"
	OperadorResta          = "-"
	OperadorMultiplicacion = "*"
	OperadorDivision       = "/"
)

// func operacionAritmetica(num1, num2 float64, operador string) float64 {
// 	var result float64
// 	switch operador {
// 	case OperadorSuma:
// 		result = num1 + num2
// 	case OperadorResta:
// 		result = num1 - num2
// 	case OperadorMultiplicacion:
// 		result = num1 * num2
// 	case OperadorDivision:
// 		result = num1 / num2
// 	}

// 	return result
// }

// func operacionAritmetica(num1, num2 float64, operador string) float64 {
// 	switch operador {
// 	case OperadorSuma:
// 		return num1 + num2
// 	case OperadorResta:
// 		return num1 - num2
// 	case OperadorMultiplicacion:
// 		return num1 * num2
// 	case OperadorDivision:
// 		if num2 != 0 {
// 			return num1 / num2
// 		}
// 	}
// 	return 0
// }

// ELLIPSIS
// puede recibir 0 o n valores
// el ellipsis DEBE ir al final
// se tratan como arrays
func sumas(nums ...float64) float64 {
	var resultado float64

	for _, value := range nums { //ignora el indice
		resultado += value
	}

	return resultado
}

func orquestadorOperacion(nums []float64, operacion func(num1, num2 float64) float64) float64 {

	var result float64

	for index, value := range nums {
		if index == 0 {
			result = 0
		} else {

			result = operacion(result, value)
		}
	}

	return result
}

// func operacionAritmetica(operador string, nums ...float64) float64 {
// 	switch operador {
// 	case OperadorSuma:
// 		return orquestadorOperacion(nums, sumas)
// 	case OperadorResta:
// 		return orquestadorOperacion(nums, resta)
// 	case OperadorMultiplicacion:
// 		return orquestadorOperacion(nums, multiplicacion)
// 	case OperadorDivision:

// 			return orquestadorOperacion(nums, division)

// 	}
// 	return 0
// }

func main() {
	// fmt.Println("Hello world")
	// a, b, c := 0, 2, -1
	// inspeccionarVariable(a)
	// inspeccionarVariable(b)
	// inspeccionarVariable(c)

	// fmt.Println(suma(13.5, 0.5))

	//fmt.Println(operacionAritmetica(2, 3, "*"))
	fmt.Println(sumas(2, 3, 4, 5))

}
