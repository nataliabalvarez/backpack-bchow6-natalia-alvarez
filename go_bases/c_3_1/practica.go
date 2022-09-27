package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id     int
	precio float64
	stock  int
}

func main() {
	//------- 1
	p1 := []byte("Id;Price;Stock\n1;5;2\n2;43;5\n3;12;0\n4;100;60")

	err := os.WriteFile("./productos.txt", p1, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Archivo creado con exito")

	//------- 2
	data, err := os.ReadFile("./productos.txt")
	if err != nil {
		panic("No se pudo abrir el archivo")
	}
	//fmt.Printf("%T", data)

	sum := 0
	for index := 0; index < len(data); index++ {
		last := 1
		if string(data[index]) == ";" {
			fmt.Print("\t")
		} else {
			fmt.Print(string(data[index]))
			if index == last+3 {
				last = index
				sum += int(data[index])
			}
		}
	}
	fmt.Printf("\n\t%d\n", sum)
}
