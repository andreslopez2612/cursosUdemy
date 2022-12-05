package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {
	andres := User{
		nombre: "andres",
		email:  "andres@gmail.com",
		activo: true,
	}

	daniela := User{
		nombre: "daniela",
		email:  "daniela@gmail.com",
		activo: true,
	}

	andresStudent := Student{
		user:   andres,
		codigo: "001aaa01",
	}

	fmt.Println(daniela)
	fmt.Println(andresStudent)

}
