package main

import (
	"curso_golang_platzi/classroomgo/carro"
	"fmt"
)

// **** STRUCTS - CLASES

func main() {
	var miCarro = carro.CarroPublic{Marca: "Renault ", Modelo: "12", Carroceria: "Berlina", Año: "1997"}
	fmt.Println(miCarro)
	fmt.Println("")

	hisCarro := carro.CarroPublic{Año: "2008", Marca: "Mazda", Modelo: "BT-50 Doble Cabina", Carroceria: "PickUp"}
	fmt.Println(hisCarro)
	fmt.Println("")

	carro.NuevoCarro("Renaul", "Sedán", "KWID", "2023")

}
