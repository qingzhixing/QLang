package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT" // add, foobar, x, y, ...
	INT   TokenType = "INT"   // 1343456

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	// 分隔符
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// 关键字
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywordMap = map[string]TokenType{
	"func": FUNCTION,
	"let":  LET,
}

func LookupIdent(ident string) TokenType {
	if token_var, exist := keywordMap[ident]; exist {
		return token_var
	}
	return IDENT
}
