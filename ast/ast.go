package ast

import "github.com/mble/monkey/token"

// Node represents the AST node
type Node interface {
	TokenLiteral() string
}

// Statement represents a non-value
// producing construct
type Statement interface {
	Node
	statementNode()
}

// Expression represents a value
// producing construct
type Expression interface {
	Node
	expressionNode()
}

// Identifier represents the IDENT
// token
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the token literal for the identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// LetStatement represents the Statement
// that tracks the identifier, token and
// expression that produces the value
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token literal for the let statement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Program represents a whole program,
// constructed of Statements
type Program struct {
	Statements []Statement
}

// TokenLiteral returns a string contianing
// the token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// ReturnStatement represents the Statement
// that tracks the token and return value
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the token literal for the return statement
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
