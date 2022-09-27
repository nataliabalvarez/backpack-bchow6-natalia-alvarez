package main

import (
	"errors"
	"fmt"
)

func calcularImpuestoSalario(salario float64) (float64, error) { //manejar el error cuando salario < 0
	var impuesto float64

	if salario < 0 {
		err = errors.New("El salario no puede ser negativo")
		return
	}

	if salario > 50000 {
		if salario > 150000 {
			impuesto = salario * 0.27
		} else {
			impuesto = salario * 0.17
		}
	}
	return impuesto

}

func calcularPromedio(notas ...float64) (float64, error) {
	var promedio float64
	var sumaNotas float64

	for _, nota := range notas {
		if nota > 0 {
			sumaNotas += nota
		} else {
			return -1, errors.New("Las notas no pueden ser de valor negativo")
		}
	}

	promedio = sumaNotas / float64(len(notas))
	return promedio, nil
}

func minsAHoras(mins float64) float64 {
	return mins / 60
}

func calcularSalario(minsTrabajados float64, categoria string) { // error por cat no existente
	switch categoria {
	case "a":
		return minsAHoras(minsTrabajados) * 1000
	case "b":
		return minsAHoras(minsTrabajados) * 1000
	case "c":
	}

}

func main() {
	//------ 1
	//fmt.Println("El impuesto es:", calcularImpuestoSalario(200000))

	//------ 2
	promedio1, error1 := calcularPromedio(1, 4, 5, -2)
	fmt.Println("Promedio con error:", promedio1, error1)
	promedio2, error2 := calcularPromedio(8, 4)
	if error2 == nil {
		fmt.Println("Promedio sin error:", promedio2)
	}

	//------ 3

}
