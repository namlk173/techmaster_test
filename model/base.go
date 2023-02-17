package model

import (
	"irishttp/helper"
	"strings"
)

type Quiz struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Answer string `json:"answer"`
	IsTrue bool   `json:"is_true"`
}

func GetQuiz(fileName string) ([]Quiz, error) {
	var quizzes []Quiz
	var answers []Answer
	var quiz Quiz

	data, err := helper.GetStringFromFile(fileName)
	if err != nil {
		return nil, err
	}

	for _, line := range data {
		isQuiz := strings.Index(line, "?") == len(line)-1
		isAnswer := strings.Index(strings.ToLower(line), "answer") == 0

		if isQuiz {
			quiz.Question = line
		} else if isAnswer {
			for i, ans := range answers {
				if ans.Answer[0] == line[len(line)-1] {
					answers[i].IsTrue = true
				}
			}
			quiz.Answers = answers
			quizzes = append(quizzes, quiz)
			answers = nil
		} else {
			answers = append(answers, Answer{Answer: line})
		}
	}

	return quizzes, nil
}
