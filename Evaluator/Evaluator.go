package Evaluator

import (
	"Taxonim/Ast"
	"time"
)

func Eval(Node Ast.INode, Env *AssertEnv, ReceivedMap map[string]interface{}) (interface{}, error) {

}

func EvalPrefixExpression(Operator string, Right interface{}) (interface{}, error) {

}

func EvalInfixExpression(Operator string, Left, Right interface{}) (interface{}, error) {

}

func EvalBangOperatorExpression(Right interface{}) (bool, error) {

}

func EvalMinusPrefixOperatorExpression(Right interface{}) (interface{}, error) {

}

func EvalFloatInfixExpression(Operator string, Left, Right float64) (interface{}, error) {

}

func EvalTimeInfixExpression(Operator string, LTime, RTime time.Time) (interface{}, error) {

}

func EvalIntegerInfixExpression(Operator string, Left, Right int64) (interface{}, error) {

}

func EvalIdentifier(Node *Ast.Identifier, Env *AssertEnv, ReceivedMap map[string]interface{}) (interface{}, error) {

}

func EvalObjectExpressions(Exps map[string]Ast.Expression, Env *AssertEnv, ReceivedMap map[string]interface{}) (map[string]interface{}, error) {

}

func EvalExpressions(Exps []Ast.Expression, Env *AssertEnv, ReceivedMap map[string]interface{}) ([]interface{}, error) {

}



func EvalCookieField(Cookie *http.Cookie, FieldName string) (interface{}, error) }{

}


func (this NotFoundError) Error() string {

}

func (this NotFoundError) Unwrap() error {

}

func (this ArgumentError) Error() string {

}

func (this ArgumentError) Unwrap() error {

}

func (this OperatorError) Error() string {

}

func (this OperatorError) Unwrap() error {

}