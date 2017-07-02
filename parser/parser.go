package parser

import (
	"github.com/mble/monkey/lexer"
	"github.com/mble/monkey/token"
)

// Parser represents the parser, containing
// a Lexer, and tracks the current token
// and the next (peek) token
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a Parser from a
// Lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens to set curToken
	// and peekToken

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
