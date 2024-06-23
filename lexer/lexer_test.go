package lexer

import (
	"QLang/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		let add = func(x, y){
			x + y;
		};
		let result = add(five, ten);
`
	fmt.Printf("input string:\n{%q}\n", input)
	testDataList := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for index, testData := range testDataList {
		token := lexer.NextToken()

		fmt.Printf("testData [ %d ] - tokenType = %q, literal = %q\n", index, token.Type, token.Literal)

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
