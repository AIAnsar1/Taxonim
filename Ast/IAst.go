package Ast

type INode interface {
	TokenLiteral() string
	String() string
}

type IStatement interface {
	INode
	statementNode()
}

type IExpression interface {
	INode
	expressionNode()
}

type IStatement interface {
	StatementNode()
	ExpressionNode()
	TokenLiteral() string
	String() string
	GetVal() interface{}
}
