package main

import "fmt"

func main() {
    // Declara un slice de enteros.
    numbers := []int{1, 2, 3, 4, 5}

    // Usa range para recorrer el slice.
    for i, num := range numbers {
        fmt.Printf("El elemento en el índice %d es %d\n", i, num)
    }

	// así recorro objetos
	// Crea un slice de Personas.
    
	type Persona struct {
		Nombre string
		Edad   int
	}
	personas := []Persona{
        {"Nico", 32},
        {"Liz", 31},
        {"Grecia", 0},
    }

    // Recorre el slice de Personas.
    for i, persona := range personas {
        fmt.Printf("Persona %d: %s, %d años\n", i, persona.Nombre, persona.Edad)
    }
}
