package lexer

import (
	"QLang/token"
)

// 解析器
type Lexer struct {
	input        string
	position     int // 所输入字符串中的当前位置（指向当前字符）
	readPosition int // 正在读取的字符（当前处理的字符的下一个字符的位置）
	character    byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func (lexer *Lexer) NextToken() token.Token {
	var token_var token.Token
	switch lexer.character {
	case '=':
		token_var = newToken(token.ASSIGN, lexer.character)
	case ';':
		token_var = newToken(token.SEMICOLON, lexer.character)
	case '(':
		token_var = newToken(token.LPAREN, lexer.character)
	case ')':
		token_var = newToken(token.RPAREN, lexer.character)
	case ',':
		token_var = newToken(token.COMMA, lexer.character)
	case '+':
		token_var = newToken(token.PLUS, lexer.character)
	case '{':
		token_var = newToken(token.LBRACE, lexer.character)
	case '}':
		token_var = newToken(token.RBRACE, lexer.character)
	case 0:
		token_var.Literal = ""
		token_var.Type = token.EOF
	}

	lexer.readChar()
	return token_var
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}
