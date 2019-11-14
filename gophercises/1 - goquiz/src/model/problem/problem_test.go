package problem

import (
	"testing"
)

func TestUser(t *testing.T) {

	question := "5+5"
	ans := "10"

	p := new(Problem)

	if p == nil {
		t.Errorf("p is not instance of Problem")
	}

	p.SetQuestion(question)
	p.SetAnswer(ans)

	if p.GetQuestion() != question {
		t.Errorf("p.question set failed, expected %v, got %v", question, p.GetQuestion())
	}
	if p.GetAnswer() != ans {
		t.Errorf("p.answer set failed, expected %v, got %v", ans, p.GetAnswer())
	}
	if !p.CheckAns("10") {
		t.Errorf("p.CheckAns(\"10\") failed, expected %v, got %v", true, bool(p.CheckAns("10")))
	}
}
