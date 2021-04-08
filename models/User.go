package models

type Tokens []string

type User struct {
	ID              uint
	email           string
	password        string
	phone           string
	address         string
	profileImageUrl string
	pushTokens      Tokens
}
