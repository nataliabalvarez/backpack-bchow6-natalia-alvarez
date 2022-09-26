package main

import "fmt"

func main() {
	fmt.Println("//------- Arrays ----") //Tamanio Estatico
	var myArray [2]int
	fmt.Println("myArray", myArray)

	var myOtherArray = [...]int{1, -2, 3, 4, 34, 123}
	fmt.Println("myOtherArray", myOtherArray)

	fmt.Println("//------- Slice ----") //Tiene un array como base. El tamanio puede cambiar dinamicamente
	slice := myOtherArray[1:3]
	fmt.Println("slice", slice)
	fmt.Println("slice len:", len(slice))
	fmt.Println("cap len:", cap(slice)) // no me da bien la capacidad

	slice = append(slice, 3, 5, -4, 1)
	fmt.Println("slice modificado:", (slice))
	fmt.Println("cap len modificado:", cap(slice))

}
