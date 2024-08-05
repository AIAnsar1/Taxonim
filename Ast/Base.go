package Ast

import (
	"Taxonim/Token"
	"go/token"
)

type ExpressionStatement struct {
	Token      Token.Token
	Expression IExpression
}

type Identifier struct {
	Token Token.Token // the token.IDENT token
	Value string
}

type Boolean struct {
	Token token.Token
	Value bool
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type FloatLiteral struct {
	Token token.Token
	Value float64
}

type NullLiteral struct {
	Token token.Token
	Value interface{}
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type ArrayLiteral struct {
	Token token.Token
	Elems []IExpression
}

type ObjectLiteral struct {
	Token token.Token
	Elems map[string]IExpression
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    IExpression
}

type InfixExpression struct {
	Token    token.Token
	Left     IExpression
	Operator string
	Right    IExpression
}

type CallExpression struct {
	Token     token.Token
	Function  IExpression
	Arguments []IExpression
}

type AssertionError struct {
	FailedAssertion string
	Received        map[string]interface{}
	WrappedError    error
}
