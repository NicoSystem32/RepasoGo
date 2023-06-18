package main

import "fmt"

func main() {
    var arr [5]int // Declara un array de 5 enteros.
    arr[0] = 1     // Asigna valores a los elementos.
    arr[1] = 2
    arr[2] = 3
    arr[3] = 4
    arr[4] = 5
    fmt.Println(arr) // Imprime: [1 2 3 4 5]

	var s []int = []int{1, 2, 3} // Declara un slice de enteros.
    s = append(s, 4) // Agrega un valor al final del slice.
    fmt.Println(s)   // Imprime: [1 2 3 4]

//Diferencias array y slices
// 	Tamaño fijo vs tamaño dinámico:
// Un array tiene un tamaño fijo. Esto significa que una vez que se declara un array con un cierto tamaño, no puedes cambiar su tamaño. Por ejemplo, si declaras un array de 5 enteros, siempre tendrá 5 enteros.
// Por otro lado, un slice no tiene un tamaño fijo. Puedes agregar o eliminar elementos de un slice, y su tamaño cambiará dinámicamente para acomodar estos cambios.
// Pasando a funciones:
// Cuando pasas un array a una función en Go, la función recibe una copia de ese array. Esto significa que si modificas el array dentro de la función, no afectará al array original.
// Sin embargo, cuando pasas un slice a una función, la función recibe una referencia al slice original. Esto significa que si modificas el slice dentro de la función, también modificarás el slice original.
// Subconjuntos:
// Es fácil crear un subconjunto de un slice (también conocido como "slicing"), mientras que no puedes hacerlo con un array. Por ejemplo, si tienes un slice con 10 elementos, puedes crear un nuevo slice que solo contenga los elementos del 5 al 8 del slice original.
// Incorporación de funciones:
// Go incluye funciones incorporadas para trabajar con slices, como append y copy. No hay funciones incorporadas similares para los arrays.
}
