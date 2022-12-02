package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool

	//Igual a
	r = a == b
	fmt.Printf("%d es igual que el segundo valor %d? %t \n", a, b, r)

	//Deferente a
	r = a != b
	fmt.Printf("%d es distinto que el segundo valor %d? %t \n", a, b, r)

	//Mayor que
	r = a > b
	fmt.Printf("%d es mayor que el segundo valor %d? %t \n", a, b, r)

	//Menor que
	r = a < b
	fmt.Printf("%d es menor que el segundo valor %d? %t \n", a, b, r)

	//Mayor o igual que
	r = a >= b
	fmt.Printf("%d es mayor que el segundo valor %d? %t \n", a, b, r)

	//Menor o igual que
	r = a <= b
	fmt.Printf("%d es menor que el segundo valor %d? %t \n", a, b, r)

}
