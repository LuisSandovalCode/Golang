package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	i := 0
	for {

		if i > 2 {
			break
		}

		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)

		i++
	}

	tiempoPasa := time.Since(inicio)

	fmt.Printf("Timpo de ejecución %s\n", tiempoPasa)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		canal <- servidor + " no se encuentra disponible"
	} else {
		canal <- servidor + " está funcionando"
	}
}
