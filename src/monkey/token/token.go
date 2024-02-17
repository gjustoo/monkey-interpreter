package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	//Identifiers + literals

	IDENT = "IDENT" // add, foobar, x, y...
	INT   = "INT"   // 134323

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
