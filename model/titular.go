package model

// Titular is a titular of account
type Titular struct {
	Nome      string `json:"nome"`
	Cpf       string `json:"cpf"`
	Profissao string `json:"profissao"`
}
