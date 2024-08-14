package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	if 1 < 2 {
		fmt.Println("1 é menor que 2")
	}

	// * Declarando variavel no if
	if x := math.Sqrt(4); x < 10 {
		fmt.Println(x)
		// short statement -> conditional (x < 10 no caso) é obrigatorio
	}
	// fmt.Println(x) -> x só existe no if

	// * Else If, Else
	if x := math.Sqrt(4); x < 10 {
		fmt.Println(x)
	} else if x > 0 {
		fmt.Println("Maior do que zero")
	} else {
		fmt.Println("Foi pro else")
	}

	if err := doError(); err != nil {
	}
}

func doError() error {
	return errors.New("Error")
}
