package parser

import (
	"github.com/KadyrPoyraz/monkeygo/ast"
	"github.com/KadyrPoyraz/monkeygo/lexer"
	"github.com/KadyrPoyraz/monkeygo/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgramm() *ast.Program {
    return nil
}
