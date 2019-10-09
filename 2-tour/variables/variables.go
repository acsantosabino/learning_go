package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// ## TIPOS
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // pseudônimo para uint8

// rune // pseudônimo para int32 representa um ponto de código Unicode

// float32 float64

// complex64 complex128

// Declaração de variaveis no pacote
var c, python, java bool

var ii, jj int = 1, 2

const Pi = 3.14

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Declaração de variaveis em uma função
	var i int
	fmt.Println(i, c, python, java)
	
	k := 3 // atribução curta dentro de uma função
	var c1, python1, java1 = true, false, "no!"
	fmt.Println(ii, jj, k, c1, python1, java1)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Convertndo tipos
	var f float64 = math.Sqrt(float64(jj*jj + k*k))
	var z uint = uint(f)
	fmt.Println(f, z)

	// Constatnte
	fmt.Println(Pi)
}