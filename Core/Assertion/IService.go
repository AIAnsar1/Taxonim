package Assertion

import "Taxonim/Core/Types"

type IAborter interface {
	AbortChan() <-chan struct{}
}

type IResultListener interface {
	Start(input <-chan *Types.ScenarioResult)
	DoneChan() <-chan struct{}
}

type IAsserter interface {
	ResultChan() <-chan TestAssertionResult
}
