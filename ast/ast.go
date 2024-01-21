package ast

import "github.com/KadyrPoyraz/monkeygo/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func NewProgram(statemets []Statement) *Program {
    return &Program{
        Statements: statemets,
    }
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

func NewIdentifier() Expression {
	return &Identifier{}
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func NewLetStatement() Statement {
	return &LetStatement{}
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (*LetStatement) statementNode() {
}
