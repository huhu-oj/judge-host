package model

type AnswerRecord struct {
	ProblemId int64
	Code string
	Input string
	Log string
	Error string
	LanguageId int
	PassNum int
	NotPassNum int
	ExecuteResultId int
}