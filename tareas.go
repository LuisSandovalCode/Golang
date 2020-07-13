package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(elemento *task) {
	t.tasks = append(t.tasks, elemento)
}

func (t *taskList) eliminarElemento(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(nuevaDescripcion string) {
	t.descripcion = nuevaDescripcion
}

func (t *task) actualizarNombre(nuevoNombre string) {
	t.nombre = nuevoNombre
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre: ", tarea.nombre)
		fmt.Println("Descripción: ", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre: ", tarea.nombre)
			fmt.Println("Descripción: ", tarea.descripcion)
		}

	}
}

func main() {

	t := &task{
		nombre:      "Completar mi curso de Go ",
		descripcion: "Completar mi curso de Go de Platzi en esta semana ",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python ",
		descripcion: "Completar mi curso de Python de Platzi en esta semana ",
	}

	t3 := &task{
		nombre:      "Completar mi curso de nodejs ",
		descripcion: "Completar mi curso de nodejs de Platzi en esta semana ",
	}

	lista := &taskList{
		tasks: []*task{
			t, t2,
		},
	}
	lista.agregarALista(t3)
	fmt.Println(lista.tasks[0])
	fmt.Println(len(lista.tasks))

	lista.imprimirLista()
	lista.tasks[0].completado = true
	lista.imprimirListaCompletados()
	//lista.eliminarElemento(1)
	// fmt.Println(len(lista.tasks))

	// for i := 0; i < len(lista.tasks); i++ {
	// 	fmt.Println("INDEX", i, "nombre", lista.tasks[i])
	// }

	// for index, tarea := range lista.tasks {
	// 	fmt.Println("INDEX", index, "nombre", tarea.nombre)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Luis"] = lista

	t4 := &task{
		nombre:      "Completar mi curso de Java ",
		descripcion: "Completar mi curso de Java de Platzi en esta semana ",
	}

	t5 := &task{
		nombre:      "Completar mi curso de c# ",
		descripcion: "Completar mi curso de c# de Platzi en esta semana ",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Alberto"] = lista2

	fmt.Println("Tareas Luis")
	mapaTareas["Luis"].imprimirLista()

	fmt.Println("Tareas Alberto")
	mapaTareas["Alberto"].imprimirLista()
}
