package go_nifi_api

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type app struct {
	username string
	password string
	host     string
	client   *http.Client
}

const api = "nifi-api"

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

func (a *app) Do(url, token, method string, data url.Values) []byte {
	var req *http.Request = nil

	switch method {
	case http.MethodGet:
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	case http.MethodPost:
		if token != "" {
			req, _ = http.NewRequest(http.MethodPost, url, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		} else {
			req, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}
	case http.MethodPut:
		req, _ = http.NewRequest(method, url, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	case http.MethodDelete:
		req, _ = http.NewRequest(method, url, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	default:
		return nil
	}

	response, _ := a.client.Do(req)

	defer response.Body.Close()

	read, _ := ioutil.ReadAll(response.Body)

	return []byte(read)
}

func NewNiFi(host string) *app {
	return &app{
		host: fmt.Sprintf("%s/%s", host, api),
	}
}
