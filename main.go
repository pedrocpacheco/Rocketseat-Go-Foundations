package main

import (
	"fmt"
	"myFirstGoProject/modulo1"
	"myFirstGoProject/pacote"
)

func main() {
	fmt.Println("hello, world!")
	pacote.PrintMinha()

	// ? Função com duas coisas
	a, b := modulo1.Swap(10, 20)
	fmt.Println(a, b)

	// ? Função retornos nomeados
	res, rem := modulo1.Dividir(10, 3)
	fmt.Println(res, rem)

	// ? High Order Function
	h := modulo1.Diminuir(3)(1)
	fmt.Println(h)

	// ? Função Variatica
	v := modulo1.Somar3(10, 10, 10)
	fmt.Println(v)

}

/*
! Criando Projeto
go mod init myFirstGoProject

! Buildando main
go build main.go

! Imports
_ "fmt"
. "fmt"
meupacote.Marshal()

! Estruturas
func Privada() {}
func publico() {}

! Variaveis
var Foo string
const Bar string = "bar"

! Estruturas
type MeuTIpo struct{}
func minhaFuncao() {}
*/
