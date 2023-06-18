package main 

import "fmt"

type Hijo struct{
	Nombre string
	Genero string
	añoNacimiento int
}

func main(){
	hijaGrecia := Hijo{
		Nombre: "Grecia Flórez Torres",
		Genero: "Femenino",
		añoNacimiento: 2023,
	}

	fmt.Println(hijaGrecia)
}