package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type escritorWeb struct{}

func (escritor escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	// respuesta, err := http.Get("http://google.com")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// e := escritorWeb{}

	// io.Copy(e, respuesta.Body)
	// fmt.Println()

	// Basic HTTP GET request
	respuesta, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
	}
	defer respuesta.Body.Close()

	// Read body from response
	body, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", body)
}
