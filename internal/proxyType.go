package internal

import "time"

type ValidProxy struct {
	ResponseTime time.Duration
	//Anonymous    bool
	ProxyType string
	Address   string
}
