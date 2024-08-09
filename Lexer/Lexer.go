package Lexer

import "Taxonim/Token"

func New(Input string) *Lexer {

}

func (this *Lexer) NewToken() {

}

func (this *Lexer) SkipWhiteSpace() {

}

func (this *Lexer) ReadChar() {

}

func (this *Lexer) PeekChar() byte {

}

func (this *Lexer) ReadIdentifier() string {

}

func (this *Lexer) ReadString() string {

}

func (this *Lexer) ReadRawString() string {

}

func (this *Lexer) ReadNumber() string {

}

func (this *Lexer) IsChAllowedInIdent(Ch byte) bool {

}

func IsLetter(Ch byte) bool {

}

func IsDigit(Ch byte) bool {

}

func NewSelfToken(TokenType Toke.TokenType, Ch byte) Token.Token {

}
