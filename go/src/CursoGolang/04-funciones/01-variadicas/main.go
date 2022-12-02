package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	fmt.Println(numeros)

	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}
	return mensaje, total
}

func main() {
	mensaje, result := sumar("Andres", 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(mensaje, result)
}
