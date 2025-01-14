package fuzz

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Fuzzer struct {
	workers, maxReq, responseT int
	blacklist                  []int
	headers                    map[string]string
	tr                         *http.Transport
}

func New(workers, maxReq, responseT int, headers []string, proxy string, blacklist []int) (Fuzzer, error) {

	var fuzzer = Fuzzer{}

	// parse headers
	h, err := pheaders(headers)
	if err != nil {
		return fuzzer, err
	}

	// parse proxy
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return fuzzer, err
	}
	return Fuzzer{
		headers:   h,
		workers:   workers,
		responseT: responseT,
		maxReq:    maxReq,
		blacklist: blacklist,
		tr:        &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
	}, nil
}

func pheaders(headers []string) (map[string]string, error) {
	rh := make(map[string]string)

	for _, h := range headers {
		p := strings.Split(h, ":")
		if len(p) < 2 {
			return nil, fmt.Errorf("%s is wrong header format", h)
		}
		rh[strings.TrimSpace(p[0])] = p[1]
	}

	return rh, nil
}
