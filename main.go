package main

import "fmt"

var nomes = []string{"Rodrigo", "Elisa"}

func listar(nomes []string) {
	fmt.Println(nomes)
}

func main() {
	listar(nomes)
}
