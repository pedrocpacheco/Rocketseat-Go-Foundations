package modulo1

import "fmt"

func DigaOi() {
	fmt.Println("oi")
}

// Um Retorno
func Somar(a, b int) int {
	return a + b
}

// Dois Retornos
func Swap(a, b int) (int, int) {
	return b, a
}

// Nomeando Retornos + Naked Return
func Dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return // ! - Naked Return -> Retorna o que ta declarado na assinatura

	/*
		? := declarando nova variavel
		? = atualizando nova variavel
	*/
}

// Função Higher Order -> Função Retorna Função
func Diminuir(a int) func(int) int {
	return func(b int) int {
		return a - b
	}
}

// Função Variatica
func Somar3(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}
