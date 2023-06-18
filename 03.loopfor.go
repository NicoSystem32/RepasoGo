package main

import "fmt"

func main(){
	fmt.Println("Nico Flórez")
	fmt.Println("im test for loop, lets go ! literally")

	//Este es el For Normalito
	for i:=0; i <= 10;i++{
		fmt.Println("Esta es la iteración nummero " + fmt.Sprint(i))
	}

	//Este es el for continued
	//Primero declaro el inicializador
	inicio := 1
	//Implementación del for el cual va a sumar los numeros de 1 hasta 10
	for; inicio <=10;{
		//re-asignación de la variable inicio
		inicio += inicio
		
	}
	fmt.Println(inicio)

	//Ejemplo de for-while ya que while es for en Go
	sumWhile := 1
	fmt.Println("Mientras... 1 < 100")
	for sumWhile < 1000 {
		sumWhile ++
		fmt.Printf("Iteración %d\n", sumWhile)
	}
	fmt.Println(fmt.Sprint(sumWhile))
}
