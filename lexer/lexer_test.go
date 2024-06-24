package lexer

import (
	"QLang/token"
	"fmt"
	"testing"
)

type DataList struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	testInput := func(input string, dataList []DataList) {
		lexer := New(input)

		for index, testData := range dataList {
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
	input_1 := `
		let five = 5;
		let ten = 10;
		let add = func(x, y){
			x + y;
		};
		let result = add(five, ten);
`
	dataList_1 := []DataList{
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
	testInput(input_1, dataList_1)

	input_2 := `
	!-/*5;
	5 < 10 > 5;
	`
	dataList_2 := []DataList{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	testInput(input_2, dataList_2)

	input_3 := `
		if(5 < 10){
			return true;
		}else{
			return false;
		}
	`

	dataList_3 := []DataList{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
	}

	testInput(input_3, dataList_3)

	input_4 := `
	10 == 10
	10 != 9
	`
	dataList_4 := []DataList{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
	}
	testInput(input_4, dataList_4)
}
