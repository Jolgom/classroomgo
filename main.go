package main

import (
	"classroomgo/dominio"
	"fmt"
)

// **** STRUCTS - CLASES

func main() {
	var miCarro = dominio.CarroPublic{Marca: "Renault ", Modelo: "12", Carroceria: "Berlina", Año: "1997"}
	fmt.Println(miCarro)
	fmt.Println("")

	hisCarro := dominio.CarroPublic{Año: "2008", Marca: "Mazda", Modelo: "BT-50 Doble Cabina", Carroceria: "PickUp"}
	fmt.Println(hisCarro)
	fmt.Println("")

	dominio.NuevoCarro("Renaul", "Sedán", "KWID", "2023")

}
