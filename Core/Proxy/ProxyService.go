package Proxy

import "net/url"

type Proxy struct {
	Strategy string
	Addr     *url.URL
	Others   map[string]interface{}
}

type SingleProxyStrategy struct {
	ProxyAddress *url.URL
}

const ProxyTypeSingle = "single"
