package pacote

import (
	"fmt"
	"myFirstGoProject/pacote/internal/foo"
)

var baz string = "hello, baz"

func print() {
	fmt.Println(baz)
}

func PrintMinha() {
	fmt.Println(foo.Minha)
}
