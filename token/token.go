package token

// Type represents the type of a token
type Type string

// Token represents the token itself, with the literal
// value and the TokenType
type Token struct {
	Type    Type
	Literal string
}

// Type constants
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// LookupIdent looks up the an indentifier to see
// if it is a keyword
func LookupIdent(ident string) Type {
	var keywords = map[string]Type{
		"fn":  FUNCTION,
		"let": LET,
	}
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
