package main

import (
	"fmt"
	"model/problem"
)

func main() {

	problems := problem.ReadProblems()
	fmt.Println(len(problems))
}
