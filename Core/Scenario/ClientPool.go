package Scenario

import (
	"Taxonim/Core/Utils"
	"net/http"
	"net/url"
)

type ClientFactoryMethod func() *http.Client
type ClientCloseMethod func(*http.Client)

var defaultFactory = func() *http.Client {
	return &http.Client{}
}

var defaultClose = func(c *http.Client) {
	c.CloseIdleConnections()
}

func NewClientPool(InitialCap int, MaxCap int, EngineMode string, Factory ClientFactoryMethod, Close ClientCloseMethod) (*Utils.Pool[*http.Client], error) {

}

func NewCookieJarRepeated() (*cookieJarRepeated, error) {

}

func (this *cookieJarRepeated) SetCookies(U *url.URL, Cookie *[]http.Cookie) {

}

func (this *cookieJarRepeated) Cookies(U url.URL) *[]http.Cookie {

}

func CreateClientFactoryMethod(Mode string, Opts ...func(http.CookieJar)) ClientFactoryMethod {

}
