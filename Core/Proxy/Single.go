package Proxy

import "net/url"

func NewProxyService(self string) (this IProxyService, err error) {

}

func init() {

}

func (this *SingleProxyStrategy) Init(self Proxy) error {

}

func (this *SingleProxyStrategy) GetAll() []*url.URL {

}

func (this *SingleProxyStrategy) GetProxy() *url.URL {

}

func (this *SingleProxyStrategy) ReportProxy(Address *url.URL, Reason string) *url.URL {

}

func (this *SingleProxyStrategy) GetProxyCountry(Address *url.URL) string {

}

func (this *SingleProxyStrategy) Done() error {

}
