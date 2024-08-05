package Report

import (
	"Taxonim/Core/Assertion"
	"Taxonim/Core/Types"
	"fmt"
	"io"
)

func init() {

}

func (this *StdoutJson) Init(Debug bool, SamplingRate int) (err error) {

}

func (this *StdoutJson) Start(Input chan *Types.ScenarioResult, AssertionResultChan <-chan Assertion.TestAssertionResult) {

}

func (this *StdoutJson) Report() {

}
func (this *StdoutJson) DoneChan() <-chan bool {

}

func (this *StdoutJson) ListenAndAggregate(Input chan *Types.ScenarioResult, AssertionResultChan <-chan *Assertion.TestAssertionResult) {

}

func (this *StdoutJson) CleanSamplingCount(SamplingCount map[uint16]map[string]int, StopSampling chan struct{}, SamplingRate int) {

}

func (this *StdoutJson) PrintInDebugMode(Input chan *Types.ScenarioResult) {

}

func PrintPretty(W io.Writer, Info any) {

}

func (this Result) MarshalJSON() ([]byte, error) {

}

func (this ScenarioStepResultSummary) MarshalJSON() ([]byte, error) {

}

func (this verboseHttpRequestInfo) MarshalJSON() ([]byte, error) {

}

var PrintJson = func(j []byte) {
	fmt.Println(string(j))
}
