package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	padraoMaisFallthrough(1)
}

// ! Break automatico

// * Fallthrough
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

// * Switch sem Condicao
func do(x int) {
	switch {
	case x == 1:
		fmt.Println(1)
	case "abc" == "foo":
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}

// * Exemplo
func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}

// * Switch que declara variavel
func matematica() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("Raiz quadrada de 4 é 2")
	default:
		fmt.Println("Algo deu errado")
	}
}

// * Varios casos na mesma linha
func isWeekend2(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

// * Switch de Tipo
func tipos(x any) {
	switch t := x.(type) {
	case int:
		fmt.Println("x é um inteiro")
	case string:
		fmt.Println("x é um string")
		takeString(t)
	default:
		fmt.Println("x é outro tipo")
	}
}

func takeString(s string) {
	fmt.Println(s)
}

// ! Switch X IFs nao muda tanto a performance
