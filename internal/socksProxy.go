package internal

import (
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func CheckSocks5Proxy(proxyList chan string, wg *sync.WaitGroup , timeout int  , valid chan ValidProxy, domain string , code int)   {


	wg.Add(1)
	defer wg.Done()

	for proxyLine := range proxyList {



		dialSocksProxy, err := proxy.SOCKS5("tcp", proxyLine, nil, proxy.Direct)
		socksTransport := &http.Transport{Dial: dialSocksProxy.Dial, DisableKeepAlives: true}

		if err != nil {
			continue
		}



		socksClient := &http.Client{
			Transport: socksTransport,
			Timeout:  time.Duration(timeout) * time.Second,
		}

		start := time.Now()
		resp, err := socksClient.Get(domain)
		if err != nil {
			continue

		}

		if resp.StatusCode != code{
			continue
		}

		_, err = ioutil.ReadAll(resp.Body)
		if err!=nil{
			continue
		}
		end := time.Since(start)

		//log.Info().Msg(string(body))


		valid <- ValidProxy{
			ResponseTime: end,
		    //	Anonymous:    true,
			ProxyType:    "socks5",
			Address:      proxyLine,
		}


		// In case you want check for anonymous proxy

		//if !strings.Contains(string(body),currentIP){
		//
		//	valid <- ValidProxy{
		//		ResponseTime: end,
		//		Anonymous:    true,
		//		ProxyType:    "socks5",
		//		Address:      proxyLine,
		//	}
		//} else {
		//
		//	valid <- ValidProxy{
		//		ResponseTime: end,
		//		Anonymous:    false,
		//		ProxyType:    "socks5",
		//		Address:      proxyLine,
		//	}
		//
		//}


		_ = resp.Body.Close()
	}


}