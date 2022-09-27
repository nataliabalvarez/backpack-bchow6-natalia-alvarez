package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Perro struct { // si algun campo esta en min, el paquete de encoding json no se va a mostrar
	Nombre string   `json:"nombre"`
	Edad   int      `json:"edad"`
	Raza   string   `json:"raza"`
	Peso   float64  `json:"peso"`
	Color  []string `json:"color"`
	Owner  Duenio   `json:"owner"`
}
type Gato struct { // si algun campo esta en min, el paquete de encoding json no se va a mostrar
	Nombre   string   `json:"nombre"`
	Edad     int      `json:"edad"`
	Raza     string   `json:"raza"`
	Peso     float64  `json:"peso"`
	Color    []string `json:"color"`
	Caracter string   `json:"caracter"`
	Owner    Duenio   `json:"owner"`
}

func (p Gato) Dormir() {
	fmt.Println("Zzzz!")
}

func (p Perro) Ladrar() {
	fmt.Println("Guau!")
}
func (p Perro) Dormir() {
	fmt.Println("Zzzz!")
}
func (p Perro) Comer() {
	fmt.Println("niom niom, que yico!")
}
func (p Perro) Saludar() {
	fmt.Println("Hola! soy", p.Nombre)
}
func (p *Perro) Renombrar(newName string) {
	p.Nombre = newName
	fmt.Println("Hola! Ahora soy", p.Nombre)
}

type Aninal interface {
	Dormir()
	Comer()
}

func NewPerro() Aninal {
	return &Perro{}
}

type Duenio struct {
	Documento  int
	Nombre     string
	NrContacto int
}

type ListaHeterogenea struct {
	Datos []interface{}
}

func main() {

	own := Duenio{
		Documento:  1234,
		Nombre:     "Nat",
		NrContacto: 123434545,
	}
	dog := Perro{
		Nombre: "firu",
		Edad:   2,
		Raza:   "salchicha",
		Peso:   5,
		Color:  []string{"negro", "blanco"},
		Owner:  own,
	}
	// fmt.Printf("Raza: %v", dog.Raza)

	// dog.Raza = "ovejero"
	// fmt.Printf("Raza: %v", dog.Raza)
	//fmt.Printf("Dog: %+v", dog)

	// dog.Owner.Documento = 1234
	// fmt.Printf("Dog: %+v", dog.Owner.Documento)
	b, err := json.Marshal(&dog) //vuelve una representacion en json en bytes y un erro
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n json.Marshal: %+v", string(b))

	dogR := reflect.TypeOf(dog)

	for i := 0; i < dogR.NumField(); i++ {
		//	field := dogR.Field(i)
		//	fmt.Printf("\nCampo: %s, tiene %s", field.Name, field.Tag.Get("json"))
	}

	//----------------- EL METODO ES CORRESPONDIENTE AL TIPO DE DATO (puede ser la extructura, obvio)
	// dog.Ladrar()
	// dog.Saludar()
	fmt.Printf("\n\nNombre antes %s\n", dog.Nombre)
	dog.Renombrar("Perrinho")
	fmt.Printf("\nNombre despues %s\n", dog.Nombre)

	perro := NewPerro()
	perro.Comer()
	perro.Dormir()

	//----------------- Interface vacia
	l := ListaHeterogenea{}
	l.Datos = append(l.Datos, 22)
	l.Datos = append(l.Datos, "Hola")
	l.Datos = append(l.Datos, 15.4)

	fmt.Println(l.Datos)
	e, ok := l.Datos[0].(string)
	fmt.Println(ok)

	fmt.Println(e) // si dio bien el tipo de dato, lo imprime, sino no

}
