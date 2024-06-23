package lexer

import (
	"QLang/token"
	"fmt"
	"strings"
)

// 解析器
type Lexer struct {
	input        string
	position     int // 所输入字符串中的当前位置（指向当前字符）
	readPosition int // 正在读取的字符（当前处理的字符的下一个字符的位置）
	character    byte
}

func (lexer *Lexer) toString() string {
	return fmt.Sprintf(
		"position:%d, readPosition:%d, character:%q",
		lexer.position,
		lexer.readPosition,
		lexer.character)
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
	fmt.Printf("after readChar(): %s\n", lexer.toString())
}

// func (lexer *Lexer) backspace() {
// 	if lexer.position == 0 {
// 		return
// 	}
// 	lexer.position--
// 	lexer.readPosition = lexer.position + 1
// 	lexer.character = lexer.input[lexer.position]
// 	fmt.Printf("Backspace(): %s\n", lexer.toString())
// }

var whiteSpaceList = " \t\n\r"

func isWhiteSpace(character byte) bool {
	return strings.Contains(whiteSpaceList, string(character))
}

func (lexer *Lexer) skipWhiteSpace() {
	for isWhiteSpace(lexer.character) {
		fmt.Printf("Skipping (val = %d) => [%q]\n", lexer.character, string(lexer.character))
		lexer.readChar()
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var token_var token.Token

	lexer.skipWhiteSpace()

	fmt.Printf("Judging character:%q\n", lexer.character)

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
	default:
		if isLetter(lexer.character) {
			token_var.Literal = lexer.readIdentifier()
			token_var.Type = token.LookupIdent(token_var.Literal)
			return token_var
		} else if isDigit(lexer.character) {
			token_var.Type = token.INT
			token_var.Literal = lexer.readNumber()
			return token_var
		} else {
			token_var = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readChar()
	return token_var
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (lexer *Lexer) readIdentifier() string {
	startPos := lexer.position
	for isLetter(lexer.character) {
		lexer.readChar()
	}
	fmt.Printf("character [%q] is not a letter\n", lexer.character)
	return lexer.input[startPos:lexer.position]
}

func isLetter(character byte) bool {
	return ('a' <= character && character <= 'z') || ('A' <= character && character <= 'Z') || character == '='
}

func (lexer *Lexer) readNumber() string {
	startPos := lexer.position
	for isDigit(lexer.character) {
		lexer.readChar()
	}
	fmt.Printf("character [%q] is not a digit\n", string(lexer.character))
	return lexer.input[startPos:lexer.position]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
