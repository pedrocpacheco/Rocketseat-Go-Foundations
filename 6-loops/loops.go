package main

import "fmt"

// ! Estrutura padrão de loop

func main() {
	var res int
	for i, a := 0, 0; i < 10; i++ {
		res++
		fmt.Println(a)
	}

	// ! Qualquer uma das 3 partes dessa estrutura são opcionais

	// * Primeira Parte Omitida
	var j int
	for ; j < 10; j++ {
		res++
	}

	// * Terceira Parte Omitida
	for j < 10 {
		res++
		j++
	}

	// * Segunda parte omitida
	for {
		res++
		j++
		break
	}
	// tudo bem que vai ser infinito, por isso o break

	// ! For Loop com Range

	// * Utilizando o Range
	arrayRange := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for range arrayRange {
		fmt.Println("dentro")
		// Só falou para fazer algo pelo tanto do range do arrayRange
	}

	// * Usando cada index do array
	for index := range arrayRange {
		fmt.Println("Index atual:", index)
	}

	// * Usando index e valor do array
	for index, valor := range arrayRange {
		fmt.Println("Index Atual:", index, "| Valor atual:", valor)
		// index e valor ainda estão dentro do 1 estrutura do for || o for pode criar direto mais de uma variavel
		// ! Quando declaramos uma variavel, ela precisa ser usada. E se eu quiser não usar o index?
	}

	// * Omitindo uma variavel e usando apenas outra
	for _, valor := range arrayRange {
		fmt.Println(valor)
		// blank identifier
	}

}
