package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

/**
Metodos
*/

//Imprimir
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

//Editar
func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

//Herencia
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Andres", 26}
	p1.imprimir()
	p1.editNombre("Brayan")
	p2 := Persona{
		nombre: "Maria",
		edad:   25,
	}
	p1.imprimir()
	p2.imprimir()

	//Herencia

	emp1 := Empleado{
		sueldo: 500,
	}
	emp1.nombre = "Jose"
	emp1.edad = 30
	emp1.imprimir()
	fmt.Println(emp1)
}
