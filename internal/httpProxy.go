package internal

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func CheckHTTPProxy(proxyList chan string, wg *sync.WaitGroup, timeout int, valid chan ValidProxy, domain string, code int) {

	//var valid []string

	wg.Add(1)
	defer wg.Done()
	for proxyLine := range proxyList {

		proxyURL, _ := url.Parse("http://" + proxyLine)
		if strings.HasPrefix(proxyLine, "http://") || strings.HasPrefix(proxyLine, "https://") {
			proxyURL, _ = url.Parse(proxyLine)
		}

		httpProxyClient := &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives: true,
				Proxy:             http.ProxyURL(proxyURL),
			},
		}

		start := time.Now()
		resp, err := httpProxyClient.Get(domain)
		if err != nil {
			continue

		}

		if resp.StatusCode != code {
			continue
		}

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}
		end := time.Since(start)

		//fmt.Println(string(body))

		valid <- ValidProxy{
			ResponseTime: end,
			//Anonymous:    false,
			ProxyType: "HTTP(s)",
			Address:   proxyLine,
		}

		// In case you want check for anonymous proxy

		//if  strings.Contains(string(body) , currentIP) {
		//	//fmt.Println(string(body))
		//	valid <- ValidProxy{
		//		ResponseTime: end,
		//		Anonymous:    false,
		//		ProxyType:    "HTTP(s)",
		//		Address:      proxyLine,
		//	}
		//	//continue
		//}else{
		//
		//	valid <- ValidProxy{
		//		ResponseTime: end,
		//		Anonymous:    true,
		//		ProxyType:    "HTTP(s)",
		//		Address:      proxyLine,
		//	}
		//
		//}

		_ = resp.Body.Close()
	}

}
