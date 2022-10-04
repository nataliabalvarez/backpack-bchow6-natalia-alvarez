package main

import (
	"fmt"
	"log"	
	"encoding/json"
)

type Product struct {
	Name string
}

func main(){
	// notebook := Product{
	// 	Name : "NOMBRE",
	// }
	// jsonData, err := json.Marshal(notebook)
	// if err != nil {
	// 	log.Fatal("error", err)
	// }
	// fmt.Printf("Producto: %s", string(jsonData))


	//-- json a go
	jsonData1 := []byte(`{"Name": "McBook"}`)
	var p Product

	if err1 := json.Unmarshal(jsonData1, &p); err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Producto: ", p)

}