package Requester

import (
	"Taxonim/Core/Types"
	"context"
	"net/url"
)

type Requester interface {
	Type() string
	Done()
}

type HttpRequesterI interface {
	Init(Ctx context.Context, Ss Types.ScenarioStep, Url *url.URL, Debug bool, Ei *Injection.EnvironmentInjector) error
}
