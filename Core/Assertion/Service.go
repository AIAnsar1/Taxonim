package Assertion

import "Taxonim/Core/Types"

func NewDefaultAssertionService() (this *DefaultAssertionService) {

}

func (this *DefaultAssertionService) Init(Assertions map[string]Types.TestAssertionOpt) chan struct{} {

}

func (this *DefaultAssertionService) GetTotalTimes() []int64 {

}

func (this *DefaultAssertionService) GetFailCount() []int64 {

}

func (this *DefaultAssertionService) Start(input <-chan *Types.ScenarioResult) {

}

func (this *DefaultAssertionService) Aggregate(self *Types.ScenarioResult) {

}

func (this *DefaultAssertionService) ApplyAssertions() {

}

func (this *DefaultAssertionService) GiveFinalResult() TestAssertionResult {

}

func (this *DefaultAssertionService) ResultChan() <-chan TestAssertionResult {

}

func (this *DefaultAssertionService) AbortChan() <-chan struct{} {

}

func (this *DefaultAssertionService) DoneChan() <-chan struct{} {

}

func (this *DefaultAssertionService) InsertSorted(self int64) {

}
