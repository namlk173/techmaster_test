package model

type Quiz struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Answer string `json:"answer"`
	IsTrue bool   `json:"is_true"`
}
