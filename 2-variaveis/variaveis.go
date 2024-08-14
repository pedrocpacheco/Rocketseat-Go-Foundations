package main

import "fmt"

var idade int

// Variaveis vazias
func Vazias() {
	var nome, sobrenome string
	fmt.Println(nome, sobrenome, idade) // "", "", 0
}

// Inicializador Variavel
func Inicializador() {
	var nome, sobrenome string = "Pedro", "Pacheco"
	fmt.Println(nome, sobrenome) // "Pedro Pacheco"
}

// Inferencia de tipo Variavel
func Inferencia() {
	var nome, sobrenome = "Pedro", "Pacheco"
	fmt.Println(nome, sobrenome) // "Pedro Pacheco"
}

// Agrupando Variavel
func Agrupando() {
	var (
		nome      = "Pedro"
		sobrenome = "Pacheco"
	)
	fmt.Println(nome, sobrenome) // "Pedro Pacheco"
}

// Sem usar var
func DoisPontos() {
	nome := "Pedro"
	sobrenome := "Pacheco"
	// sobrenome = Pacheco não pode
	fmt.Println(nome, sobrenome)
}

// Declarando em escopo de pacote
var idade2 = 30

// Omitindo em escopo de pacote
var idade3 = 40

// Na mesma linha tipos diferentes
var foo, bar, zoo = "foo", 50, false

// Na mesma linha que não pode -> sem valor
// ! var baz, qux, int, string => nao pode
