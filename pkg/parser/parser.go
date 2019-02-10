package lexer

import (
	"io"
)

type TokenType int

const (
	TokenTypeEmpty TokenType = iota
	TokenLParen
	TokenRParen
	TokenPlus
	TokenMinus
)

func next(reader io.RuneScanner) error {
	return nil
}

func GetNextToken(reader io.Reader) {
	reader.Read()
}
