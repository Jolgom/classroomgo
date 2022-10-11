package main

import "fmt"

// **** STRUCTS - CLASES
// Se hace por fuera del main.

type carro struct {
	carroceria, marca, modelo, año string
}

func main() {
	miCarro := carro{marca: "Renault ", modelo: "12", carroceria: "Berlina", año: "1997"}
	fmt.Println(miCarro)
	fmt.Println("")

	hisCarro := carro{año: "2008", marca: "Mazda", modelo: "BT-50 Doble Cabina", carroceria: "PickUp"}
	fmt.Println(hisCarro)
	fmt.Println("")

	// OTRA MANERA
	var herCarro carro
	herCarro.marca = "Renaul"
	herCarro.carroceria = "Sedán"
	herCarro.modelo = "KWID"
	fmt.Println(herCarro)
	fmt.Println("")

	fmt.Println(herCarro.marca, herCarro.carroceria)
	fmt.Println(miCarro.marca, miCarro.carroceria)
	fmt.Println(hisCarro.marca, hisCarro.carroceria)
}
