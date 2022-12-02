package main

import "fmt"

func main() {
	users := make(map[string]string)

	users["Andres"] = "andres@gmail.com"
	users["Maria"] = "maria@gmail.com"

	if correo, ok := users["Jose"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
