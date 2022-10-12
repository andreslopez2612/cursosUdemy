package main

import "fmt"

func main() {

	var numero1 int

	fmt.Print("Por favor ingrese el valor de venta del producto: ")
	fmt.Scanln(&numero1)

	result := (numero1 * 18) / 100

	fmt.Println("El IGV es:", result)

	result = numero1 + result

	fmt.Println("El precio de venta es:", result)

}
