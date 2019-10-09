package main

import "fmt"

// Delarando a primeira função
func add(x int, y int) int {
	return x + y
}

// Declaração curta de parametros, todos os parametros podem ter o tipo ocultado exceto do ultimo
func add_short(x, y int) int{
	return x + y
}

// Retorno de multiplos resultados
func swap(x, y string) (string, string) {
	return y, x
}

// Resultados nomeados e retorno limpo
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	
	fmt.Println(add_short(42, 13))
	
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	fmt.Println(split(17))
}