package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	MinimoNoImponible = 150000
)

// ---------------- Ejercicio 1
type SalaryCustomError struct {
	Msg string
}

func (e *SalaryCustomError) Error() string {
	return e.Msg
}

func DebePagarImpuesto(salary int) (err error) {
	if salary < MinimoNoImponible {
		err = &SalaryCustomError{
			Msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
		return
	}
	return nil
}

// ---------------- Ejercicio 2
var SalaryCustomError2 = errors.New("error: el salario ingresado no alcanza el mínimo imponible")

func DebePagarImpuesto2(salary int) error {
	if salary < MinimoNoImponible {

		return SalaryCustomError2
	}
	return nil
}

// ---------------- Ejercicio 3
func DebePagarImpuesto3(salary int) error {
	if salary < MinimoNoImponible {

		err := fmt.Errorf("error: el minimo no imponible es de 150000 y el salario ingresado es de %d", salary)
		return err

	}
	return nil
}

func main() {
	fmt.Println("\n\n**---------------------- Ejercicio 1")
	salary := 170000
	err1 := DebePagarImpuesto(salary)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")

	fmt.Println("\n\n**---------------------- Ejercicio 2")
	salary2 := 180000
	err2 := DebePagarImpuesto2(salary2)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")

	fmt.Println("\n\n**---------------------- Ejercicio 3")
	salary3 := 140000
	err3 := DebePagarImpuesto3(salary3)
	if err3 != nil {
		fmt.Println(err3)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")

}
