package Ast

func (this *ExpressionStatement) StatementNode()       {}
func (this *ExpressionStatement) TokenLiteral() string {}
func (this *ExpressionStatement) String() string {

}

func (this *Identifier) ExpressionNode()      {}
func (this *Identifier) TokenLiteral() string {}
func (this *Identifier) String() string {

}

func (this *Boolean) ExpressionNode() {}
func (this *Boolean) TokenLiteral()   {}
func (this *Boolean) String() string {

}
func (this *Boolean) GetVal() interface{} {

}

func (this *IntegerLiteral) ExpressionNode()      {}
func (this *IntegerLiteral) TokenLiteral() string {}
func (this *IntegerLiteral) String() string {

}
func (this *IntegerLiteral) GetVal() interface{} {

}

func (this *FloatLiteral) ExpressionNode()      {}
func (this *FloatLiteral) TokenLiteral() string {}
func (this *FloatLiteral) String() string {

}
func (this *FloatLiteral) GetVal() interface{} {

}

func (this *NullLiteral) ExpressionNode()      {}
func (this *NullLiteral) TokenLiteral() string {}
func (this *NullLiteral) String() string {

}
func (this *NullLiteral) GetVal() interface{} {

}

func (this *StringLiteral) ExpressionNode()      {}
func (this *StringLiteral) TokenLiteral() string {}
func (this *StringLiteral) String() string {

}
func (this *StringLiteral) GetVal() interface{} {

}

func (this *ArrayLiteral) ExpressionNode()      {}
func (this *ArrayLiteral) TokenLiteral() string {}
func (this *ArrayLiteral) String() string {

}
func (this *ObjectLiteral) ExpressionNode()      {}
func (this *ObjectLiteral) TokenLiteral() string {}
func (this *ObjectLiteral) String() string {

}

func (this *PrefixExpression) ExpressionNode()      {}
func (this *PrefixExpression) TokenLiteral() string {}
func (this *PrefixExpression) String() string {

}

func (this *InfixExpression) ExpressionNode()      {}
func (this *InfixExpression) TokenLiteral() string {}
func (this *InfixExpression) String() string {

}

func (this *CallExpression) ExpressionNode()      {}
func (this *CallExpression) TokenLiteral() string {}
func (this *CallExpression) String() string {

}
