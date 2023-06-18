package main
import "fmt"

func main(){
	fmt.Println("Hello world Nico con Golang")
	// ejemplos de variables y su formato 
	numero := 42
	decimal := 3.14
	cadena := "Hola, mundo!"
	booleano := true
	caracter := 'A'
	arreglo := []int{1, 2, 3}
	mapa := map[string]int{"clave": 123}
	puntero := &numero
	estructura := struct {
		Nombre string
		Edad   int
	}{Nombre: "Juan", Edad: 30}

	fmt.Printf("Número: %d\n", numero)
	fmt.Printf("Decimal: %f\n", decimal)
	fmt.Printf("Cadena: %s\n", cadena)
	fmt.Printf("Booleano: %t\n", booleano)
	fmt.Printf("Caracter: %c\n", caracter)
	fmt.Printf("Arreglo: %v\n", arreglo)
	fmt.Printf("Mapa: %v\n", mapa)
	fmt.Printf("Puntero: %p\n", puntero)
	fmt.Printf("Estructura: %+v\n", estructura)
}