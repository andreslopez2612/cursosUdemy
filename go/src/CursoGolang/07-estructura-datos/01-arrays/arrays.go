package main

import "fmt"

func main() {
	//arrays
	var numeros [5]int

	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[3] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[1])

	//arrays con valores
	nombres := [3]string{"Brayan", "Andres", "Lopez"}

	fmt.Println(nombres)
	colores := [...]string{
		"Rojo",
		"Negro",
		"Verde",
		"Azul",
	}

	fmt.Println(colores, len(colores))

	//indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}

	fmt.Println(monedas, len(monedas))

	subMonedas := monedas[0:3]

	fmt.Println(subMonedas, len(subMonedas))
}
