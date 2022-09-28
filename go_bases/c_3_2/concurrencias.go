package main

import (
	"fmt"
	"time"
)

func procesar(i int, ch chan int) {
	fmt.Println(i, "- inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "- termina")
	ch <- i
}

func main() {
	canal := make(chan int)
	now := time.Now()
	fmt.Println("Inicia el programa")
	for i := 0; i < 10; i++ {
		go procesar(i, canal) // las go routines se ejecutan en "segundo plano" y el main termina antes que ellas
		// el tiempo maximo es el tiempo de 1 proceso
	}
	for i := 0; i < 10; i++ {
		lectura := <-canal
		fmt.Println("Lectura de canal:", lectura)

	}
	//time.Sleep(2000 * time.Millisecond)
	fmt.Println("Termina el programa")
	fmt.Printf("el programa demoro %d msegs", time.Now().Sub(now).Milliseconds())

}
