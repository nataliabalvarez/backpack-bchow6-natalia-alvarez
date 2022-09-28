package main

import (
	"errors"
	"fmt"
	"os"
)

type MyCustomError struct { // los errores customizados siempre deben implementar la interface error
	StatusCode int
	Msg        string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.StatusCode, e.Msg)
}

func ObtainError(status int) (code int, err error) {
	if status >= 400 {
		code = 500
		err = &MyCustomError{
			StatusCode: code,
			Msg:        "Algo salio mal",
		}
		return
	}
	return 200, nil
}

var error1 = errors.New("error 1")

func getError() error {
	return fmt.Errorf("extra info %w", error1)
}

type ErrType2 struct{}

func (e ErrType2) Error() string {
	return "soy el error 2"
}

func main() {
	err1 := fmt.Errorf("First error :/") //permite generar un erro

	//continua la ejecicion
	//ocurre el error 2

	err2 := fmt.Errorf("Second error: %w", err1) //permite generar un erro
	err3 := errors.New("Third error")

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)

	//--- errores custom
	status, err := ObtainError(300)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status: %d, funciona!", status) //no toma > 400

	//---- new error
	// statusCode := 404
	// if statusCode > 399 {
	// 	error = errors.New("La peticion ha fallado")
	// 	fmt.Println(error) // :(
	// 	os.Exit(1)

	// }
	// fmt.Printf("Status: %d, funciona!", statusCode)

	//----- Package Errors implementa funciones para manipular errores
	// new crea errores
	//--- AS compprueba si un error es de un tipo especifico, por los errores personalizados
	// error1 := &MyCustomError{
	// 	StatusCode: 502,
	// 	Msg:        "soy el error 1",
	// }

	// error2 := MyCustomError{
	// 	StatusCode: 404,
	// 	Msg:        "soy el error 2",
	// }

	// bothErrorsAreEquals := errors.As(error1, error2)
	// fmt.Println(bothErrorsAreEquals)

	//-- IS comprueba un error con un valor
	errorr := getError()
	coincidence := errors.Is(errorr, error1)
	fmt.Println("\n", coincidence)

	//-- Unwrap si un error1 contiene a un error2, con Unwrap, obtenemos el error2

	errType2 := ErrType2{}
	errType1 := fmt.Errorf("soy err 1, contengo a %w", errType2)

	fmt.Println("Unrap error tipo1", errors.Unwrap(errType1))
	fmt.Println("Unrap error tipo2", errors.Unwrap(errType2))

}
