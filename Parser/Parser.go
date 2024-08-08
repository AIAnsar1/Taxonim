package Parser

import (
	"Taxonim/Ast"
	"Taxonim/Lexer"
	"Taxonim/Token"
)

func New(this *Lexer.Lexer) *Parser {

}

func (this *Parser) NexToken() {

}

func (this *Parser) CurTokenIs(self *Token.TokenType) bool {

}

func (this *Parser) PeekTokenIs(self *Token.TokenType) bool {

}

func (this *Parser) ExpectPeek(self *Token.TokenType) bool {

}

func (this *Parser) Errors() []string {

}

func (this *Parser) PeekError(self *Token.TokenType) {

}

func (this *Parser) ParseExpressionStatement() *Ast.ExpressionStatement {

}

func (this *Parser) ParseExpression(Precedence int) *Ast.IExpression {

}

func (this *Parser) PeekPrecedence() int {

}

func (this *Parser) CurPrecedence() int {

}

func (this *Parser) ParseIdentifier() int {

}

func (this *Parser) ParseIdentifierLiteral() Ast.IExpression {

}

func (this *Parser) ParseFloatLiteral() Ast.IExpression {

}

func (this *Parser) ParseStringLiteral() Ast.IExpression {

}

func (this *Parser) ParsePrefixExpression() Ast.IExpression {

}

func (this *Parser) ParseInfixExpression(left Ast.IExpression) Ast.IExpression {

}

func (this *Parser) ParseBoolean() Ast.IExpression {

}

func (this *Parser) ParseNull() Ast.IExpression {

}

func (this *Parser) ParseGroupExpression() Ast.IExpression {

}

func (this *Parser) ParseObjectLiteral() Ast.IExpression {

}

func (this *Parser) ParseArrayLiteral() Ast.IExpression {

}

func (this *Parser) ParseCallExpression() Ast.IExpression {

}

func (this *Parser) ParseCallArguments() []Ast.IExpression {

}

func (this *Parser) ParseArrayElements() []Ast.IExpression {

}

func (this *Parser) ParseObjectElements() map[string]Ast.IExpression {

}
