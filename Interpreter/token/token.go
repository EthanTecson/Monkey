// Ethan Tecson
// 2/8/2025

package token

type TokenType string

type Token struct {
	Type		TokenType
	Litearl		string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifers + literals
	IDENT = "IDENT" // add, foobar, x, y ...
	INT = "INT" // 123456789
	
	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)