package tipos

import (
	"fmt"
	"strconv"
)

// bool

// * Inteiros
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr

// ! Int -> Inteiros Positivos e Negativos
// ! Uint -> Inteiros apenas Positivos

// ? uintptr -> apenas coisas unsafe

// byte
// ? mesma coisa que uint8

// rune
// ? mesma coisa que int32

// float32 float64

// complex64 complex128
// ? Apenas isso é 128

// string

func main() {
	// somar("a", "b")
	var b uint8 = 10
	takeByte(b)
}

func somar(a, b int) int {
	return a + b
}

func takeByte(b byte) {
	fmt.Println(b)
}

func conversao() {
	var i int
	i = 65 // i := 65
	f := float64(i)
	s := string(i) // A maiosculo -> 65 na unicode utf8
	fmt.Println(f, s)
	sCorreto := strconv.FormatInt(int64(i), 10)
	fmt.Println(sCorreto)
}

// ! string(int) converte para runa e utf no primeiro ""
