package Lexer




type ILexer interface {
	NewToken()
	SkipWhiteSpace() 	
	ReadChar()
	PeekChar() byte
	ReadIdentifier() string 
	ReadString() string)
	ReadRawString() string
	ReadNumber() string
    IsChAllowedInIdent(Ch byte) bool
}