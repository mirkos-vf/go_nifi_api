package go_nifi_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go_nifi_api/types"
)

func (a *app) ControllerServices(id, token string, method string) *types.ControllerServiceEntity {

	controllerServiceEntity := types.ControllerServiceEntity{}

	url := fmt.Sprintf("%s/%s/%s", a.host, "controller-services", id)
	bytes, _ := a.Do(url, token, http.MethodGet, nil)
	_ = json.Unmarshal(bytes, &controllerServiceEntity)

	return &controllerServiceEntity
}
