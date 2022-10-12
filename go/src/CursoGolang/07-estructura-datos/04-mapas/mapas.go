package main

import "fmt"

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	dias[7] = "SABADO"

	fmt.Println(dias)

	delete(dias, 1)

	fmt.Println(dias, len(dias))

	//nuevo mapa
	estudiantes := make(map[string][]int)

	estudiantes["Andres"] = []int{13, 14, 15}
	estudiantes["Lopez"] = []int{16, 17, 18}

	fmt.Println(estudiantes, len(estudiantes))
	fmt.Println(estudiantes["Andres"], len(estudiantes))
	fmt.Println(estudiantes["Andres"][1], len(estudiantes))
}
