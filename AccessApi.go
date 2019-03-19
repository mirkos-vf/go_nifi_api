package go_nifi_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (a *app) AccessToken(username, password string) (token string, err error) {
	a.makeClient()

	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", a.host, "nifi-api/access/token"), strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}

	response, err := a.client.Do(req)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
