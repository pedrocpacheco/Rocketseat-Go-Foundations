package main

import "fmt"

func main() {
	arrayZerado := [4]int{}  // Array precisa ter tamanho constante
	fmt.Println(arrayZerado) // [0, 0, 0]

	arrayInicializado := [4]int{1, 2, 3}
	fmt.Println(arrayInicializado) // [1, 2, 3, 0]

	indexEspecifico := [3]int{1: 2} // index: valor
	fmt.Println(indexEspecifico)

	// ! Só pode passar constante como valor de array

	// * var naoConstante = 10
	// * arrayComNaoConstante := [naoConstante]int{} -> Tamanho de index não aceito

	const constante = 10
	arrayComConstante := [constante]int{} // * Tamanho de index aceito

	fmt.Println(arrayComConstante)

	// ! Array nunca pode mudar de tamanho

}
