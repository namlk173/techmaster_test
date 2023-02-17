package model

import (
	"fmt"
	"testing"
)

func TestGetQuiz(t *testing.T) {
	fileName := "../data/cauhoimau.txt"
	quizzes, err := GetQuiz(fileName)
	if err != nil {
		t.Errorf("Can't found file data: %s", fileName)
	}
	fmt.Println(quizzes)
}
