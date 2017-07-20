package ast

import (
	"bytes"

	"github.com/mble/monkey/token"
)

// Node represents the AST node
type Node interface {
	TokenLiteral() string
	String() string
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

// String returns the identifier value as a string
func (i *Identifier) String() string {
	return i.Value
}

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

// String returns a string the contains the token and value
// of the let statement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
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

// String returns a string that contains the return
// value of each statements String() method
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String returns the token and return value as a string
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents the Statement
// that tracks the token and expression
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the token literal for the expression statement
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns the expression as a string
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// IntegerLiteral represents integer literal
// expressions
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the integer literal expression
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns the integer literal expression as a string
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
