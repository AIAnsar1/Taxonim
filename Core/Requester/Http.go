package Requester

import (
	"Taxonim/Core/Types"
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"
)

func (this *HttpRequester) Init(Ctx context.Context, S Types.ScenarioStep, ProxyAddress *url.URL, Debug bool, Ei *Injection.EnvironmentInjector) (err error) {

}
func (this *HttpRequester) Done() {

}
func (this *HttpRequester) Send(Client *http.Client, Envs map[string]interface{}) (Res *Types.ScenarioStepResult) {

}
func (this *HttpRequester) PrepareReq(Envs map[string]interface{}, Trace *httptrace.ClientTrace) (*http.Request, error) {

}
func (this *HttpRequester) InitTransport() *http.Transport {

}
func (this *HttpRequester) UpdateTransport(Tranc *http.Transport) {

}
func (this *HttpRequester) InitTLSConfig() *tls.Config {

}
func (this *HttpRequester) InitRequestInstance() (err error) {

}
func (this *HttpRequester) Type() string {

}
func (this *HttpRequester) ApplyAssertions(AssertEnv *Evaluator.AssertEnv) (bool []Types.FailedAssertion) {

}
func (this *HttpRequester) CaptureEnvironmentVariables(Header http.Header, ResponseBody []byte, Cookies map[string]*http.Cookie, ExtractedVars map[string]interface{}) map[string]string {

}
func (this *Duration) SetResStartTime(Time time.Time) {

}
func (this *Duration) SetServerProcessStart(Time time.Time) {

}
func (this *Duration) SetDNSDuration(Time time.Time) {

}
func (this *Duration) DetDNSDuration(Time time.Time) {

}
func (this *Duration) SetTLSDuration(Time time.Time) {

}
func (this *Duration) GetTLSDuration(Time time.Time) {

}
func (this *Duration) SetConnDuration(Time time.Time) {

}
func (this *Duration) GetConnDuration(Time time.Time) {

}
func (this *Duration) SetReqDuration(Time time.Time) {

}
func (this *Duration) GetReqDuration(Time time.Time) {

}
func (this *Duration) SetServerProcessDuration(Time time.Time) {

}
func (this *Duration) GetServerProcessDuration(Time time.Time) {

}
func (this *Duration) SetResDuration(Time time.Time) {

}
func (this *Duration) GetResDuration(Time time.Time) {

}
func (this *Duration) IotalDuration(Time time.Time) {

}

func ConcatEnvs(Input map[string]interface{}, Output map[string]interface{}) {

}

func ConcatHeaders(Input map[string]interface{}, Output map[string]interface{}) {

}

func FetchErrType(err error) Types.RequestError {

}

func NewTrace(Duration *Duration, ProxyAddress *url.URL, HeadersByClient map[string][]string) *httptrace.ClientTrace

func (this *Duration) close() {

}
