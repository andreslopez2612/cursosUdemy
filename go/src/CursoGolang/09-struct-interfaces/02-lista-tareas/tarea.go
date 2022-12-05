package main

import "fmt"

//Lista de tareas
type TaskList struct {
	task []*Task
}

//

func (tl *TaskList) appendTask(t *Task) {
	tl.task = append(tl.task, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.task = append(tl.task[:index], tl.task[index+1:]...)
}

//Tareas
type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar curso de GO este mes",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de HTML",
		desc:      "Completar curso de HTML esta semana",
		completed: true,
	}

	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)

	// t1.toPrint()
	// t2.toPrint()

	fmt.Println(lista)

	t3 := Task{
		name:      "Curso de CSS",
		desc:      "Completar curso de CSS esta semana",
		completed: true,
	}

	lista.appendTask(&t3)
	fmt.Println(lista)

	lista.removeTask(1)

	for _, task := range lista.task {
		fmt.Println(task.name)
		fmt.Printf("Nombre: %s \nDescripcion: %s", task.name, task.desc)
	}

}
