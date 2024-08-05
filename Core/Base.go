package Core

import (
	"Taxonim/Core/Assertion"
	"Taxonim/Core/Proxy"
	"Taxonim/Core/Report"
	"Taxonim/Core/Scenario"
	"Taxonim/Core/Types"
	"context"
	"sync"
)

type Engine struct {
	Hammer           Types.Hammer
	ProxyService     Proxy.IProxyService
	ReportService    Report.ReportService
	ScenarioService  *Scenario.ScenarioService
	Aborter          Assertion.IAborter
	Asserter         Assertion.IAsserter
	ResListener      Assertion.IResultListener
	TickCounter      int
	ReqCountArr      []int
	Wg               sync.WaitGroup
	ResultReportChan chan *Types.ScenarioResult
	ResultAssertChan chan *Types.ScenarioResult
	AbortChan        <-chan struct{}
	TestSuccess      bool
	Ctx              context.Context
}

type EngineServices struct {
	Aborter     Assertion.IAborter
	Asserter    Assertion.IAsserter
	ResListener Assertion.IResultListener
	ProxyServ   Proxy.IProxyService
	ReportServ  Report.ReportService
}
