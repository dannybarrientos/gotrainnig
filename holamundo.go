package main

import (
	fmt "fmt"
)

func main() {
	var usuario string
	fmt.Print("Escribe tu nombre a continuacion: ")
	fmt.Scanf("%s", &usuario)
	fmt.Printf("Hola Mundoo, %s !\n", usuario)
}
