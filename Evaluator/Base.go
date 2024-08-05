package Evaluator

import "net/http"

type AssertEnv struct {
	StatusCode    int64
	ResponseSize  int64
	ResponseTime  int64 // in ms
	Body          string
	Headers       http.Header
	Variables     map[string]interface{}
	Cookies       map[string]*http.Cookie // cookies sent by the server, name -> cookie
	TotalTime     []int64                 // in ms
	FailCount     int
	FailCountPerc float64 // should be in range [0,1]
}

type NotFoundError struct {
	Source       string
	WrappedError error
}

type ArgumentError struct {
	Message      string
	WrappedError error
}

type OperatorError struct {
	Message      string
	WrappedError error
}
