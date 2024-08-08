package Ast

import "Taxonim/Evaluator"

func (this AssertionError) Error() string {

}

func (this AssertionError) Rule() string {

}

func (this AssertionError) WReceived() map[string]interface{} {

}

func (this AssertionError) Unwrap() error {

}

func Assert(Input string, Env *Evaluator.AssertEnv) (bool, error) {

}
