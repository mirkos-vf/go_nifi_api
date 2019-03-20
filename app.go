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

func (a *app) Do(url, token, method string, data url.Values) ([]byte, error) {
	var req *http.Request = nil
	var err error

	req, err = http.NewRequest(method, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	if token == "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	response, _ := a.client.Do(req)

	defer response.Body.Close()

	read, _ := ioutil.ReadAll(response.Body)

	return []byte(read), nil
}

func NewNiFi(host string) *app {
	return &app{
		host: fmt.Sprintf("%s/%s", host, api),
	}
}
