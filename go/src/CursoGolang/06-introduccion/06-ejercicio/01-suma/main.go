package main

import "fmt"

func main() {

	var numero1 int
	var numero2 int

	fmt.Print("Por favor ingrese el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Print("Por favor ingrese el segundo numero: ")
	fmt.Scanln(&numero2)

	result := numero1 + numero2

	fmt.Println("El resultado es:", result)

}
