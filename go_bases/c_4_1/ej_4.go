package main

import "fmt"

var (
	MinimoNoImponible = 150000
	HorasMinimas      = 80
)

type SalaryCustomError struct {
	Msg string
}

func (e *SalaryCustomError) Error() string {
	return fmt.Sprintf("- %v", e.Msg)
}

type HsCustomError struct {
	Msg string
}

func (e *HsCustomError) Error() string {
	return fmt.Sprintf("- %v", e.Msg)
}
func CalcularSalario(horasTrabajadas int, valorPorHora float64) (salario float64, err error) {
	if horasTrabajadas < 80 && horasTrabajadas > 0 {
		err = &HsCustomError{
			// puedo usar erros.new
			Msg: "error: el trabajador no puede haber trabajado menos de 80 hs mensuales",
		}
		return
	} else if horasTrabajadas < 0 {
		err = &HsCustomError{
			Msg: "error: el valor mensual trabajado no puede ser negativo",
		}
	}
	return 0, nil
}

func main() {

}
