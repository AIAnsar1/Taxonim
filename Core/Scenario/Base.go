package Scenario

import (
	"Taxonim/Core/Injection"
	"Taxonim/Core/Requester"
	"Taxonim/Core/Types"
	"Taxonim/Core/Utils"
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

type ScenarioService struct {
	Clients     map[*url.URL][]ScenarioItemRequester
	CPool       *Utils.Pool[*http.Client]
	Scenario    Types.Scenario
	Ctx         context.Context
	ClientMutex sync.Mutex
	Debug       bool
	EngineMode  string
	Ei          *Injection.EnvironmentInjector
	IterIndex   int64
}

type ScenarioOpts struct {
	Debug                  bool
	IterationCount         int
	MaxConcurrentIterCount int
	EngineMode             string
	InitialCookies         []*http.Cookie
}

type ScenarioItemRequester struct {
	scenarioItemID uint16
	sleeper        Sleeper
	requester      Requester.Requester
}

type RangeSleep struct {
	min int
	max int
}

type DurationSleep struct {
	duration int
}

type cookieJarRepeated struct {
	defaultCookieJar *cookiejar.Jar
	firstIterPassed  bool
}
