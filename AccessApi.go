//Get controller configuration, Manage the cluster, Create reporting tasks
package go_nifi_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts.
// The header, the body, and the signature. The expiration of the token is a contained within the body.
// The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
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

// The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued.
// It is used as a query parameter name 'access_token'.
func (a *app) UiExtensionToken(username, password string) (token string, err error) {
	// TODO: func UiExtensionToken

	return "", nil
}
