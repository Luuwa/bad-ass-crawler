package crawler

import "math/rand"

type Proxy struct {
	proxyAdress string
}
type ProxyList struct {
	proxyList []Proxy
}

func (p *ProxyList) random() Proxy {
	if len(p.proxyList) >= 1 {
		randomIndex := rand.Intn(len(p.proxyList))
		randomProxy := p.proxyList[randomIndex]
		return randomProxy
	} else {
		return Proxy{proxyAdress: ""}
	}

}
func createProxyList(proxyAdress []string) ProxyList {
	proxies := []Proxy{}
	for _, proxy := range proxyAdress {
		proxies = append(proxies, Proxy{proxyAdress: proxy})
	}
	ProxyList := ProxyList{proxyList: proxies}
	return ProxyList
}
