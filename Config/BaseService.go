package Config

import "Taxonim/Core/Types"

type ConfigReader interface {
	Init([]byte) error
	CreateHammer() (Types.Hammer, error)
}

type TimeRunCount []struct {
	Duration int `json:"duration"`
	Count    int `json:"count"`
}

type Auth struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type MultipartFormData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
	Src   string `json:"src"`
}

type RegexCaptureConf struct {
	Exp *string `json:"exp"`
	No  int     `json:"matchNo"`
}
type CapturePath struct {
	JsonPath   *string           `json:"json_path"`
	XPath      *string           `json:"xpath"`
	XpathHtml  *string           `json:"xpath_html"`
	RegExp     *RegexCaptureConf `json:"regexp"`
	From       string            `json:"from"`
	CookieName *string           `json:"cookie_name"`
	HeaderKey  *string           `json:"header_key"`
}

type Step struct {
	Id               uint16                 `json:"id"`
	Name             string                 `json:"name"`
	Url              string                 `json:"url"`
	Auth             Auth                   `json:"auth"`
	Method           string                 `json:"method"`
	Headers          map[string]string      `json:"headers"`
	Payload          string                 `json:"payload"`
	PayloadFile      string                 `json:"payload_file"`
	PayloadMultipart []MultipartFormData    `json:"payload_multipart"`
	Timeout          int                    `json:"timeout"`
	Sleep            string                 `json:"sleep"`
	Others           map[string]interface{} `json:"others"`
	CertPath         string                 `json:"cert_path"`
	CertKeyPath      string                 `json:"cert_key_path"`
	CaptureEnv       map[string]CapturePath `json:"capture_env"`
	Assertions       []string               `json:"assertion"`
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

type JsonReader struct {
	ReqCount     *int                   `json:"request_count"`
	IterCount    *int                   `json:"iteration_count"`
	LoadType     string                 `json:"load_type"`
	Duration     int                    `json:"duration"`
	Assertions   []TestAssertion        `json:"success_criterias"`
	TimeRunCount TimeRunCount           `json:"manual_load"`
	Steps        []Step                 `json:"steps"`
	Output       string                 `json:"output"`
	Proxy        string                 `json:"proxy"`
	Envs         map[string]interface{} `json:"env"`
	Data         map[string]CsvConf     `json:"data"`
	Debug        bool                   `json:"debug"`
	SamplingRate *int                   `json:"sampling_rate"`
	EngineMode   string                 `json:"engine_mode"`
	Cookies      CookieConf             `json:"cookie_jar"`
}

type CookieConf struct {
	Cookies []CustomCookie `json:"cookies"`
	Enabled bool           `json:"enabled"`
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

type TestAssertion struct {
	Rule  string `json:"rule"`
	Abort bool   `json:"abort"`
	Delay int    `json:"delay"`
}

type RemoteMultipartError struct {
	Message      string
	wrappedError error
}

var AvailableConfigReader = make(map[string]ConfigReader)

const ConfigTypeJson = "jsonReader"
