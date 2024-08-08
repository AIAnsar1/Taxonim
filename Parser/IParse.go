package Parser

import (
	"Taxonim/Ast"
	"Taxonim/Token"

)

type IParse interface {
	NexToken()
	CurTokenIs(self *Token.TokenType) bool
	PeekTokenIs(self *Token.TokenType) bool
	ExpectPeek(self *Token.TokenType) bool
	Errors() []string
	PeekError(self *Token.TokenType)
	ParseExpressionStatement() *Ast.ExpressionStatement
	ParseExpression(Precedence int) *Ast.IExpression
	PeekPrecedence() int
	CurPrecedence() int
	ParseIdentifier() int
	ParseIdentifierLiteral() Ast.IExpression
	ParseFloatLiteral() Ast.IExpression
	ParseStringLiteral() Ast.IExpression
	ParsePrefixExpression() Ast.IExpression
	ParseInfixExpression(left Ast.IExpression) Ast.IExpression
	ParseBoolean() Ast.IExpression
	ParseNull() Ast.IExpression
	ParseGroupExpression() Ast.IExpression
	ParseObjectLiteral() Ast.IExpression
	ParseArrayLiteral() Ast.IExpression
	ParseCallExpression() Ast.IExpression
	ParseCallArguments() []Ast.IExpression
	ParseArrayElements() []Ast.IExpression
	ParseObjectElements() map[string]Ast.IExpression
}
