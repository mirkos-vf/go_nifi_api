package nifi_api

import (
	"crypto/tls"
	"net/http"
)

const (
	api_url = "https://vm-pd-nifi-1.dh.rt.ru:8080/nifi-api"
)

type app struct {
	Token string
}

type client struct {
	client *http.Client
}

func (c *client) makeClient() {
	// Create New http Transport
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	// Create Http Client
	c.client = &http.Client{Transport: transCfg}
}

func NewNiFi() *app {
	return &app{}
}
