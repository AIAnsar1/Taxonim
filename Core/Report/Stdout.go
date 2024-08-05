package Report

import (
	"Taxonim/Core/Assertion"
	"Taxonim/Core/Types"
	"io"
)

func init() {

}

func (this *Stdout) Init(Debug bool, SamplingRate int) (err error) {

}

func (this *Stdout) Start(Input chan *Types.ScenarioResult, AssertionResultChan <-chan Assertion.TestAssertionResult) {

}

func (this *Stdout) CleanSamplingCount(SamplingCount map[uint16]map[string]int, SamplingStop chan struct{}, SamplingRate int) {

}

func (this *Stdout) Report() {

}
func (this *Stdout) DoneChan() <-chan bool {

}

func (this *Stdout) RealTimePrintStart() {

}

func (this *Stdout) LiveResultPrint() {

}

func (this *Stdout) RealTimePrintStop() {

}

func (this *Stdout) PrintInDebugMode(Input chan *Types.ScenarioResult) {

}

func PrintBody(W io.Writer, ContentType string, BOdy interface{}) {

}

func (this *Stdout) PrintDetails() {

}

func Deduplicate(Values []interface{}) []interface{} {

}
