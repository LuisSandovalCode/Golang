package main

import "fmt"

type perro struct{}

type pajaro struct{}

type pez struct{}

func (perro) caminiar() string {
	return "soy un perro y camino"
}

func (pez) nadar() string {
	return "soy un pez y nado"
}

func (pajaro) volar() string {
	return "soy un pajaro y vuelo"
}

func moverPerro(p perro) {
	fmt.Println(p.caminiar())
}

func moverPez(p pez) {
	fmt.Println(p.nadar())
}
func moverPajaro(p pajaro) {
	fmt.Println(p.volar())
}

func main() {
	perro := perro{}

	moverPerro(perro)

	pajaro := pajaro{}

	moverPajaro(pajaro)

	pez := pez{}

	moverPez(pez)
}
