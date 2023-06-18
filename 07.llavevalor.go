package main

import "fmt"

func main(){
	//voy a crear un diccionario
	miDiccionario:= make(map[string]int)
	miDiccionario["Nico Flórez"] = 32
	miDiccionario["Liz Torres"] = 31
	miDiccionario["Grecia Flórez"] = 0

	fmt.Println(miDiccionario)// imprime todo el diccionario 
	//map[Grecia Flórez:0 Liz Torres:31 Nico Flórez:32]

	//Si quiero imprimir uno especial por ejemplo edad de Grecita

	g := miDiccionario["Grecia"] //se imprime 0 
	fmt.Println(g)
	// otro diccionario
	cargosDic := make(map[string] string)
	cargosDic["Developer"] = "Nico Flórez"
	cargosDic["Qa Analist"] = "Liz Torres"
	cargosDic["Directora de desarrollo"] = "Grecia Flórez"
	fmt.Println("Todos los empleados")	
	fmt.Println(cargosDic)
	fmt.Println(cargosDic["Developer"])
	//Recorriendo el diccionario
	for i, e := range cargosDic {
		fmt.Println(i + ":", e)
	}

}