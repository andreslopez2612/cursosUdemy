package main

import "fmt"

func main() {
	//slicen
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	//agregar datos

	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	numeros = append(numeros, 6)

	fmt.Println(numeros, len(numeros))

	//sub slicen

	subSlicen := numeros[:2]

	numeros[0] = 100

	fmt.Println(subSlicen, len(subSlicen))
	fmt.Println(numeros, len(numeros))

	//punteros
	//longitud
	//capacidad

	meses := []string{
		"Enero",
		"Febrero",
		"Marzo",
	}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
