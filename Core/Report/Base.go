package Report

import (
	"Taxonim/Core/Assertion"
	"Taxonim/Core/Types"
	"github.com/fatih/color"
	"sync"
	"time"
)

type ReportService interface {
	DoneChan() <-chan bool
	Init(debug bool, samplingRate int) error
	Start(input chan *Types.ScenarioResult, assertionResultChan <-chan Assertion.TestAssertionResult)
}

type Result struct {
	TestStatus           string                                `json:"test_status"`
	TestFailedAssertions []Assertion.FailedRule                `json:"failed_criterias"`
	SuccessCount         int64                                 `json:"success_count"`
	ServerFailedCount    int64                                 `json:"server_fail_count"`
	AssertionFailCount   int64                                 `json:"assertion_fail_count"`
	AvgDuration          float32                               `json:"avg_duration"`
	StepResults          map[uint16]*ScenarioStepResultSummary `json:"steps"`
}

type AssertionErrVerbose struct {
	Count      int64                  `json:"count"`
	Conditions map[string]*AssertInfo `json:"conditions"`
}

type ServerErrVerbose struct {
	Count   int64          `json:"count"`
	Reasons map[string]int `json:"reasons"`
}

type FailVerbose struct {
	Count              int64               `json:"count"`
	AssertionErrorDist AssertionErrVerbose `json:"assertions"`
	ServerErrorDist    ServerErrVerbose    `json:"server"`
}

type ScenarioStepResultSummary struct {
	Name           string             `json:"name"`
	StatusCodeDist map[int]int        `json:"status_code_dist"`
	Fail           FailVerbose        `json:"fail"`
	Durations      map[string]float32 `json:"durations"`
	SuccessCount   int64              `json:"success_count"`
}

type AssertInfo struct {
	Count    int
	Received map[string][]interface{}
	Reason   string
}

type verboseRequest struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

type verboseResponse struct {
	StatusCode   int               `json:"status_code"`
	Headers      map[string]string `json:"headers"`
	Body         interface{}       `json:"body"`
	ResponseTime int64             `json:"response_time"`
}

type verboseHttpRequestInfo struct {
	StepId           uint16                  `json:"step_id"`
	StepName         string                  `json:"step_name"`
	Request          verboseRequest          `json:"request"`
	Response         verboseResponse         `json:"response"`
	Envs             map[string]interface{}  `json:"envs"`
	TestData         map[string]interface{}  `json:"test_data"`
	FailedCaptures   map[string]string       `json:"failed_captures"`
	FailedAssertions []Types.FailedAssertion `json:"failed_assertions"`
	Error            string                  `json:"error"`
}

type Stdout struct {
	DoneChan     chan bool
	Result       *Result
	PrintTicker  *time.Ticker
	Mu           sync.Mutex
	Debug        bool
	SamplingRate int
}

type Duration struct {
	Name     string
	Duration float32
	Order    int
}

type StdoutJson struct {
	doneChan     chan bool
	result       *Result
	debug        bool
	samplingRate int
	mu           sync.Mutex
}

type alias struct {
	StepId           uint16                  `json:"step_id"`
	StepName         string                  `json:"step_name"`
	Envs             map[string]interface{}  `json:"envs"`
	TestData         map[string]interface{}  `json:"test_data"`
	FailedCaptures   map[string]string       `json:"failed_captures"`
	FailedAssertions []Types.FailedAssertion `json:"failed_assertions"`
	Request          struct {
		Url     string            `json:"url"`
		Method  string            `json:"method"`
		Headers map[string]string `json:"headers"`
		Body    interface{}       `json:"body"`
	} `json:"request"`
	Response verboseResponse `json:"response"`
}

type Report Result
type ItemReport ScenarioStepResultSummary

var (
	White                   = color.New(color.FgHiWhite).SprintFunc()
	Blue                    = color.New(color.FgHiBlue).SprintFunc()
	Green                   = color.New(color.FgHiGreen).SprintFunc()
	Yellow                  = color.New(color.FgHiYellow).SprintFunc()
	Red                     = color.New(color.FgHiRed).SprintFunc()
	RealTimePrintInterval   = time.Duration(1500) * time.Millisecond
	AvailableOutputServices = make(map[string]ReportService)
)

var keyToStr = map[string]Duration{
	"dnsDuration":           {Name: "DNS", Order: 1},
	"connDuration":          {Name: "Connection", Order: 2},
	"tlsDuration":           {Name: "TLS", Order: 3},
	"reqDuration":           {Name: "Request Write", Order: 4},
	"serverProcessDuration": {Name: "Server Processing", Order: 5},
	"resDuration":           {Name: "Response Read", Order: 6},
	"duration":              {Name: "Total", Order: 7},
}
var strKeyToJsonKey = map[string]string{
	"dnsDuration":           "dns",
	"connDuration":          "connection",
	"tlsDuration":           "tls",
	"reqDuration":           "request_write",
	"serverProcessDuration": "server_processing",
	"resDuration":           "response_read",
	"duration":              "total",
}

const (
	OutputTypeStdoutJson = "stdout-json"
	OutputTypeStdout     = "stdout"
)
