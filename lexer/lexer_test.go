package lexer

import (
	"QLang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	testDataList := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for index, testData := range testDataList {
		token := lexer.NextToken()

		if token.Type != testData.expectedType {
			t.Fatalf("testData [ %d ] - tokenType wrong.Expected = %q, got = %q",
				index, testData.expectedType, token.Type)
		}

		if token.Literal != testData.expectedLiteral {
			t.Fatalf("testData [ %d ] - literal wrong.Expected = %q, got = %q",
				index, testData.expectedLiteral, token.Literal)
		}
	}
}
