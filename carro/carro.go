package carro

import "fmt"

// *** Acceso PUBLICO O PRIVADO

// Para generar una clas con acceso publico a las variables (Public)
// Se debe Nombrar el struct (CLASE) con la letra capital en MAYUSCULA
// Y comentarlo usanto el mismo nombre del struct

// CarroPublic - CLASE CARRO, CON ACCESO PUBLICO.
type CarroPublic struct {
	Carroceria, Marca, Modelo, Año string
}

type carroPrivate struct {
	carroceria, marca, modelo, año string
}

// NuevoCarro
func NuevoCarro(Carroceria, Marca, Modelo, Año string) {
	var automovil carroPrivate
	automovil.marca = Marca
	automovil.carroceria = Carroceria
	automovil.modelo = Modelo
	automovil.año = Año
	fmt.Println(automovil)
}
