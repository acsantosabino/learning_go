package main

import (
	"flag"
	"fmt"
	"model/problem"
	"time"
)

var asw string
var score, total int

func quiz(filepath string, ch chan int) {
	problems := problem.ReadProblems(filepath)
	score = 0
	total = len(problems)

	for i, p := range problems {
		fmt.Printf("%d. ", i+1)
		p.Ask()
		fmt.Scanf("%s", &asw)
		if p.CheckAns(asw) {
			fmt.Println("Good")
			score++
		} else {
			fmt.Println("Wrong")
		}
	}
	ch <- 1
}

func main() {
	ch := make(chan int)
	filePtr := flag.String("file", "", "File path of quiz problems")
	timePtr := flag.Int("time", 30, "Time to complete the quiz in seconds")
	flag.Parse()

	fmt.Println("You have ", *timePtr, "s to complete the quiz. Press the ENTER key to start")
	fmt.Scanln()
	go quiz(*filePtr, ch)

	select {
	case <-ch:
		fmt.Println("End!")
	case <-time.After(time.Duration(*timePtr) * time.Second):
		fmt.Println("\nTime is over!")
	}
	fmt.Println("Your Score: ", score, "of ", total)
}
