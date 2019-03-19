package go_nifi_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go_nifi_api/types"
)

func (a *app) Controllerservices(id, token string) *types.ControllerServiceEntity {

	controllerServiceEntity := types.ControllerServiceEntity{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", a.host, "controller-services", id), nil)
	req.Header.Set("Content-Type", "*/*")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, _ := a.client.Do(req)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	bytes := []byte(data)
	_ = json.Unmarshal(bytes, &controllerServiceEntity)

	return &controllerServiceEntity
}
