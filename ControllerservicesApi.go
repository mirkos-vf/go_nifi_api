package go_nifi_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"go_nifi_api/types"
)

func (a *app) Controllerservices(id, token string) *types.ControllerServiceEntity {

	controllerServiceEntity := types.ControllerServiceEntity{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", a.host, "nifi-api/controller-services", id), nil)
	req.Header.Add("Content-Type", "*/*")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, _ := a.client.Do(req)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	bytes := []byte(data)
	_ = json.Unmarshal(bytes, &controllerServiceEntity)

	return &controllerServiceEntity
}

func (a *app) ControllerservicesPut(id, token string, entity *types.ControllerServiceEntity) *types.ControllerServiceEntity {

	controllerServiceEntity := types.ControllerServiceEntity{}

	body, _ := json.Marshal(entity)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", a.host, "nifi-api/controller-services", id), strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, _ := a.client.Do(req)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	bytes := []byte(data)
	_ = json.Unmarshal(bytes, &controllerServiceEntity)

	return &controllerServiceEntity
}
