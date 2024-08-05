package Core

import "time"

type ICore interface {
	IsTestFailed() bool
	Init() (err error)
	Start() string
	RunWorkers(C int)
	RunWorker(ScenarioStartTime time.Time)
	RunAssertionsInEngine() bool
	Stop()
	GetMaxConcurrentIterCount() int
	InitReqCountArr()
	CreateManualReqCountArr()
	CreateLinearReqCountArr()
	CreateIncrementalReqCountArr()
	CreateWavedReqCountArr()
	CreateLinearDistArr()
	CreateIncrementalDistArr()
}
