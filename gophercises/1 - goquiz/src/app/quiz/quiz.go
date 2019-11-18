package main

import (
	"fmt"
	"model/problem"
)

func main() {

	var asw string
	var score int
	problems := problem.ReadProblems()

	for i, p := range problems {
		score = i
		fmt.Printf("%d. ", i)
		p.Ask()
		fmt.Scanf("%s", &asw)
		if p.CheckAns(asw) {
			fmt.Println("Good")
		} else {
			fmt.Println("Wrong")
			break
		}
	}

	fmt.Println("Game Over!\n Your Score: ", score, "of ", len(problems)-1)
}
