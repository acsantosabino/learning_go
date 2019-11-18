package problem

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Problem struct {
	question string
	answer   string
}

var defealtPath = "./data/problems.csv"

func (p *Problem) SetQuestion(q string) {
	p.question = q
}
func (p *Problem) GetQuestion() string {
	return p.question
}

func (p *Problem) SetAnswer(a string) {
	p.answer = a
}
func (p *Problem) GetAnswer() string {
	return p.answer
}

func (p *Problem) Ask() {
	fmt.Printf("What is %v = ? ", p.question)
}

func (p *Problem) CheckAns(ans string) bool {
	return p.answer == ans
}

func ReadProblems(path string) (list []Problem) {

	if path == "" {
		path = defealtPath
	}

	f, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		q := new(Problem)
		q.SetQuestion(record[0])
		q.SetAnswer(record[1])
		list = append(list, *q)
	}

	return list
}
