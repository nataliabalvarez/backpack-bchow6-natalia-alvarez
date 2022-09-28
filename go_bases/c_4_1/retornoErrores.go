package main

import (
	"errors"
	"fmt"
	"os"
)

func division(nro1, nro2 float64) (result float64, err error) {
	if nro2 == 0 {
		err = errors.New("no puede divdir por cero")
		return
	}
	result = nro1 / nro2

	return
}
func main() {
	result, err := division(18, 0)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("resultado: %v \n", result)

}
