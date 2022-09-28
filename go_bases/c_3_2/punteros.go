package main

import "fmt"

func incrementar(puntero *int) {
	*puntero = 20
}

type Elemento struct {
	Valor int
}

func (c *Elemento) Actualizar(valor int) {
	c.Valor = valor
}

func main() {
	//var puntero *int
	//p2:=new (int)
	var numero int = 10
	p3 := numero

	fmt.Printf("La direccion es %p\n", &p3)
	fmt.Printf("La direccion es %p\n", &numero)

	//*p3 = 20
	//numero = 20

	incrementar(&numero)

	fmt.Printf("el valor es %d\n", numero)
	c := Elemento{Valor: 2}
	fmt.Printf("el valor es %d\n", c.Valor)

	c.Actualizar(35)
	fmt.Printf("el valor es %d\n", c.Valor)

}
