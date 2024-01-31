package ast

import (
	"testing"

	"github.com/KadyrPoyraz/monkeygo/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

    expectedString := "let myVar = anotherVar;"
	if program.String() != expectedString {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
