package Types

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	ErrorProxy          = "[ ERROR ]: Proxy"
	ErrorConn           = "[ ERROR ]: Connection"
	ErrorUnkown         = "[ ERROR ]: Unknown"
	ErrorIntented       = "[ ERROR ]: Intented"
	ErrorDns            = "[ ERROR ]: DNS"
	ErrorParse          = "[ ERROR ]: Parse"
	ErrorAddr           = "[ ERROR ]: Address"
	ErrorInvalidRequest = "[ ERROR ]: Invalid Request"
	ReasonProxyFailed   = "[ ERROR ]: Proxy Connection Refused"
	ReasonProxyTimeout  = "[ ERROR ]: Proxy Timeout"
	ReasonConnTimeout   = "[ ERROR ]: Connection Timeout"
	ReasonReadTimeout   = "[ ERROR ]: Read Timeout"
	ReasonConnRefused   = "[ ERROR ]: Connection Refused"
	ReasonCtxCanceled   = "[ ERROR ]: Context Canceled"

	// Constants of the Load Types
	LoadTypeLinear      = "Linear"
	LoadTypeIncremental = "Incremental"
	LoadTypeWaved       = "Waved"

	EngineModeDistinctUser = "Distinct-user"
	EngineModeRepeatedUser = "Repeated-user"
	EngineModeTaxonim      = "Taxonim"

	// Default Values
	DefaultIterCount     = 100
	DefaultLoadType      = LoadTypeLinear
	DefaultDuration      = 10
	DefaultTimeout       = 5
	DefaultMethod        = http.MethodGet
	DefaultOutputType    = "stdout"
	DefaultSamplingCount = 3
	DefaultSingleMode    = true

	ProtocolHTTP                           = "HTTP"
	ProtocolHTTPS                          = "HTTPS"
	AuthHttpBasic                          = "Basic"
	maxSleep                               = 90000
	EnvironmentVariableRegexStr            = `{{[a-zA-Z$][a-zA-Z0-9_().-]*}}`
	EnvironmentVariableNameStr             = `^[a-zA-Z][a-zA-Z0-9_-]*$`
	Header                      SourceType = "Header"
	Body                        SourceType = "Body"
	Cookie                      SourceType = "Cookies"
)

type RequestError struct {
	Type   string
	Reason string
}

type ScenarioValidationError struct {
	Message      string
	WrappedError error
}

type EnvironmentNotDefinedError struct {
	Message      string
	WrappedError error
}

type CaptureConfigError struct {
	Message      string
	WrappedError error
}

type FailedAssertion struct {
	Rule     string
	Received map[string]interface{}
	Reason   string
}

type TestAssertionOpt struct {
	Abort bool
	Delay int
}

type TimeRunCount []struct {
	Duration int
	Count    int
}

type Tag struct {
	Tag  string `json:"tag"`
	Type string `json:"type"`
}

type CsvConf struct {
	Path          string         `json:"path"`
	Delimiter     string         `json:"delimiter"`
	SkipFirstLine bool           `json:"skip_first_line"`
	Vars          map[string]Tag `json:"vars"`
	SkipEmptyLine bool           `json:"skip_empty_line"`
	AllowQuota    bool           `json:"allow_quota"`
	Order         string         `json:"order"`
}

type CustomCookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Expires  string `json:"expires"`
	MaxAge   int    `json:"max_age"`
	HttpOnly bool   `json:"http_only"`
	Secure   bool   `json:"secure"`
	Raw      string `json:"raw"`
}

type Hammer struct {
	IterationCount    int
	LoadType          string
	TestDuration      int
	TimeRunCountMap   TimeRunCount
	Scenario          Scenario
	Proxy             Proxy.Proxy
	ReportDestination string
	Others            map[string]interface{}
	Debug             bool
	SamplingRate      int
	EngineMode        string
	TestDataConf      map[string]CsvConf
	Cookies           []CustomCookie
	CookiesEnabled    bool
	Assertions        map[string]TestAssertionOpt
	SingleMode        bool
}

type ScenarioResult struct {
	// First request start time for the Scenario
	StartTime time.Time

	ProxyAddr   *url.URL
	StepResults []*ScenarioStepResult

	// Dynamic field for extra data needs in response object consumers.
	Others map[string]interface{}
}

type ScenarioStepResult struct {
	StepID           uint16
	StepName         string
	RequestID        uuid.UUID
	StatusCode       int
	RequestTime      time.Time
	Duration         time.Duration
	ContentLength    int64
	Err              RequestError
	Url              string
	Method           string
	ReqHeaders       http.Header
	ReqBody          []byte
	RespHeaders      http.Header
	RespBody         []byte
	Custom           map[string]interface{}
	UsableEnvs       map[string]interface{}
	ExtractedEnvs    map[string]interface{}
	FailedCaptures   map[string]string
	FailedAssertions []FailedAssertion
}

type SourceType string

type ScenarioStep struct {
	ID            uint16
	Name          string
	Method        string
	Auth          Auth
	Cert          tls.Certificate
	CertPool      *x509.CertPool
	Headers       map[string]string
	Payload       string
	URL           string
	Timeout       int
	Sleep         string
	Custom        map[string]interface{}
	EnvsToCapture []EnvCaptureConf
	Assertions    []string
}

type RegexCaptureConf struct {
	Exp *string `json:"exp"`
	No  int     `json:"matchNo"`
}

type EnvCaptureConf struct {
	JsonPath   *string           `json:"json_path"`
	Xpath      *string           `json:"xpath"`
	XpathHtml  *string           `json:"xpath_html"`
	RegExp     *RegexCaptureConf `json:"regexp"`
	Name       string            `json:"as"`
	From       SourceType        `json:"from"`
	Key        *string           `json:"header_key"`
	CookieName *string           `json:"cookie_name"`
}

type CsvData struct {
	Rows   []map[string]interface{}
	Random bool
}

type Auth struct {
	Type     string
	Username string
	Password string
}

var LoadTypes = [...]string{LoadTypeLinear, LoadTypeIncremental, LoadTypeWaved}
var EngineModes = [...]string{EngineModeTaxonim, EngineModeDistinctUser, EngineModeRepeatedUser}
var SupportedProtocols = [...]string{ProtocolHTTP, ProtocolHTTPS}
var supportedProtocolMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodHead, http.MethodOptions}
var supportedAuthentications = []string{AuthHttpBasic}
var envVarRegexp *regexp.Regexp
var envVarNameRegexp *regexp.Regexp
