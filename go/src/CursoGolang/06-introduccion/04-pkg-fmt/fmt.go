package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Print(hola)
	fmt.Print(mundo)

	nombre := "Alex"
	edad := 26

	//fmt.Println(hola)
	//fmt.Println(mundo)

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	fmt.Printf("Nombre : %T \n", nombre)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ootro nombre: ", nombre)
}
