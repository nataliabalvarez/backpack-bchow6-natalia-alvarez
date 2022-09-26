package main

import "fmt"

func main() {
	fmt.Println("//------- Maps ----")
	var mapa = map[string]int{"Nat": 31, "Lucas": 16}
	fmt.Println(mapa)
	fmt.Println("Lucas: ", mapa["Lucas"])
	x, ok := mapa["Nat"]
	fmt.Println(x, ok)
	//_, ok := mapa["Nat"]
	//fmt.Println(ok)
}
