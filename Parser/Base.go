package Parser

import (
	"Taxonim/Ast"
	"Taxonim/Lexer"
	"Taxonim/Token"
	"go/token"
)

type Parser struct {
	L              *Lexer.Lexer
	Errors         []string
	CurToken       token.Token
	PeekToken      token.Token
	PrefixParseFns map[Token.TokenType]PrefixParseFn
	InfixParseFns  map[Token.TokenType]InfixParseFn
}

type (
	PrefixParseFn func() Ast.IExpression
	InfixParseFn  func(Ast.IExpression) Ast.IExpression
)

var precedences = map[Token.TokenType]int{
	Token.EQ:       EQUALS,
	Token.NOT_EQ:   EQUALS,
	Token.LT:       LESSGREATER,
	Token.GT:       LESSGREATER,
	Token.PLUS:     SUM,
	Token.MINUS:    SUM,
	Token.SLASH:    PRODUCT,
	Token.ASTERISK: PRODUCT,
	Token.LBRACKET: ARRAYDEFINE,
	Token.LBRACE:   OBJECTDEFINE,
	Token.LPAREN:   CALL,
	Token.AND:      ANDOR,
	Token.OR:       ANDOR,
}

const (
	_ int = iota
	LOWEST
	ANDOR
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	ARRAYDEFINE
	OBJECTDEFINE
	CALL
)
