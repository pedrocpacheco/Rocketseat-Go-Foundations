package main

import "fmt"

func main() {
	padraoMaisFallthrough(1)
}

func padraoMaisFallthrough(x int) {
	switch x {
	case 1:
		fmt.Println("Um")
		fallthrough // -> Passa pro proximo
	case 2:
		fmt.Println("Dois")
	default:
		fmt.Println("Outro")
	}
}
