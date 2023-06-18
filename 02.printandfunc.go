package main

import "fmt"

func main(){
	fmt.Println(SumaDosNumeros(3, 5,"Este es el mensaje 1", "Este es el mensaje 2",true)) 
	fmt.Println("Pruebas")
	var tipoRune = 'N'
	fmt.Println(string(tipoRune))// si no se parsea imprime el código unicode
}

func SumaDosNumeros(a , b int, message1, message2 string, boleano bool) (int,string, string, bool){
	result := a + b
	return result, message1, message2, boleano
}