package go_nifi_api

import (
	"crypto/tls"
	"net/http"
)

type app struct {
	// Token    string
	username string

	password string

	host string

	client *http.Client
}

func (a *app) makeClient() {
	// Create New http Transport
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	// Create Http Client
	a.client = &http.Client{
		Transport: transCfg,
	}
}

func NewNiFi(host string) *app {
	return &app{
		host: host,
	}
}
