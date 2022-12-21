package mensajes

import "fmt"

func Hola() { //publica
	fmt.Println("Hola desde el paquete")
}

const mensaje = "Hola desde la constante"

func funcionPrivada() { //privada
	fmt.Println("Hola desde la funcion privada")
}

func Mensaje() {
	fmt.Println(mensaje)
	funcionPrivada()
}
