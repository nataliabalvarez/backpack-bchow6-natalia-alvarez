package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	structure "github.com/nataliabalvarez/backpack-bchow6-natalia-alvarez/go_web/structures"
)


type Product struct {
	Id           int
	Name         string
	Color        string
	Price        float64
	Stock        int
	Code         string
	Published    bool
	CreationDate string
}

type Products struct {
	Products []Product
}

func main() {
	router := gin.Default()

	router.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola, Nat!",
		})
	})

	//------------------- ejercicio 3
	router.GET("/productos", GetAll)

	router.Run()
}

func GetAll(ctx *gin.Context) {

	var prods Products

	prods.ReadJson("productos.json")
	
	ctx.JSON(200, prods)
}

func (prods *Products) ReadJson(name string) {

	jsonprods, err := os.Open(name)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonprods)

	if err1 := json.Unmarshal(byteValue, &prods); err1 != nil {
		log.Fatal(err1)
	}

}
