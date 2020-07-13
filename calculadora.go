package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(entrada string, operador string) {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		break
	case "-":
		fmt.Println(operador1 - operador2)
		break
	case "*":
		fmt.Println(operador1 * operador2)
		break
	case "/":
		fmt.Println(operador1 / operador2)
		break
	default:
		fmt.Println(operador, "No esta soportado")
		break
	}

}

func parsear(entrada string) int {
	num, _ := strconv.Atoi(entrada)

	return num
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}

	c.operate(entrada, operador)
}
