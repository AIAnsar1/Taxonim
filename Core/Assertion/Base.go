package Assertion

import (
	"Taxonim/Core/Types"
	"sync"
)

type DefaultAssertionService struct {
	Assertions map[string]Types.TestAssertionOpt // Rule -> Opts
	AbortChan  <-chan struct{}
	DoneChan   <-chan struct{}
	ResChan    chan TestAssertionResult
	AssertEnv  *Evaluator.AssertEnv
	AbortTick  map[string]int // rule -> tickIndex
	IterCount  int
	Mu         sync.Mutex
}

type TestAssertionResult struct {
	Fail        bool         `json:"fail"`
	Aborted     bool         `json:"aborted"`
	FailedRules []FailedRule `json:"failed_rules"`
}

type FailedRule struct {
	Rule        string                 `json:"rule"`
	ReceivedMap map[string]interface{} `json:"received"`
}

var TickerInterval = 100
