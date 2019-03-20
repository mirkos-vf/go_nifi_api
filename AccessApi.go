//Get controller configuration, Manage the cluster, Create reporting tasks
package go_nifi_api

import (
	"fmt"
	"net/http"
	"net/url"
)

// The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts.
// The header, the body, and the signature. The expiration of the token is a contained within the body.
// The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
func (a *app) AccessToken(username, password string) (token string, err error) {
	a.makeClient()

	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	urls := fmt.Sprintf("%s/%s", a.host, "access/token")
	bytes, err := a.Do(urls, "", http.MethodPost, formData)

	return string(bytes), err
}
