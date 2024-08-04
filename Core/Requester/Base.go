package Requester

import (
	"Taxonim/Core/Types"
	"context"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

type HttpRequester struct {
	Ctx                  context.Context
	ProxyAddr            *url.URL
	Packet               Types.ScenarioStep
	Client               *http.Client
	Request              *http.Request
	Ei                   *Injection.EnvironmentInjector
	ContainsDynamicField map[string]bool
	ContainsEnvVar       map[string]bool
	Debug                bool
	DynamicRgx           *regexp.Regexp
	EnvRgx               *regexp.Regexp
}

type Duration struct {
	DnsDuration                time.Duration
	ConnectionDuration         time.Duration
	TlsDuration                time.Duration
	RequestDuration            time.Duration
	ResponseDuration           time.Duration
	ServerProcessDuration      time.Duration
	ResStart                   time.Time
	ResStartCh                 chan time.Time
	ResStartChClosed           bool
	ServerProcessDurCh         chan time.Duration
	ServerProcessDurChClosed   bool
	ServerProcessStartCh       chan time.Time
	ServerProcessStartChClosed bool
	ResDurCh                   chan time.Duration
	ResDurChClosed             bool
	Mu                         sync.Mutex
	ChMu                       sync.Mutex
	GetChLock                  sync.Mutex
}
