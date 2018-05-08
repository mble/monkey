package ast

import (
	"bytes"
	"strings"

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

// Boolean represents the TRUE and FALSE
// tokens
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral returns the token literal for the boolean
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String returns the boolean value as a string
func (b *Boolean) String() string { return b.Token.Literal }

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

// BlockStatement represents blocks, which are like
// mini-Programs, consisting of many Statements
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns the token literal for the block statement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// String returns the string representation of the block statement
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// PrefixExpression represents prefix expressions
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns the token literal for prefix expression
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns the prefix expression as a parentheses-wrapped
// string
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents infix expressions
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns the token literal for infix expression
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns the infix expression as a parentheses-wrapped
// string
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// IfExpression represents if control-flow expressions,
// where the consequence and alternative are block statements
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns the token literal for the if expression
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String returns the if statement as a string
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// FunctionLiteral represents function literals,
// definted with the "fn" keyword
// It is constructed of a token, a slice of identifier pointers
// and a block statement for the function body
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the token literal for the function literal
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// String returns the function literal as a string
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
