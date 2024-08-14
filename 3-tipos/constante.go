package tipos

import "fmt"

func constante() {
	// * Caso mostrando declaração tipo
	// ! Existe Erro !!
	// const x int = 10
	// takeInt32(x) // -> não posso usar pq é int não int32

	// * Caso não mostrando tipo
	// ? Não existe erro
	const x = 10
	takeInt32(x) // Funciona!
	// ! Dribla de leve os tipos em go
	// ! Untyped Const -> se adequa segundo o contexto

	takeInt64(x)
}

// ? Nem todo tipo pode ser constante
// ! Apenas Caracteres
// ! Booleanos
// ! Númericos
// ! String
// * Qualquer outro tipo não pode ser const
// * Não pode usar := para criar constante

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

// constante literal
