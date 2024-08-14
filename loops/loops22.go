package main

import "fmt"

func loops22() {

	// ! Interando sobre um inteiro
	for range 10 {
		fmt.Println("dentro")
	}

	for valor := range 10 {
		fmt.Println(valor)
		// Só pode declarar uma variavel -> não pode ter for v, a := range 10 {}
	}

	// ? Condição de corrida -> Duas escritas ao mesmo tempo
	// i := i

}
