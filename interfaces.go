package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct{}

type pajaro struct{}

type pez struct{}

func (perro) mover() string {
	return "soy un perro y camino"
}

func (pez) mover() string {
	return "soy un pez y nado"
}

func (pajaro) mover() string {
	return "soy un pajaro y vuelo"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	perro := perro{}

	moverAnimal(perro)

	pajaro := pajaro{}

	moverAnimal(pajaro)

	pez := pez{}

	moverAnimal(pez)
}
