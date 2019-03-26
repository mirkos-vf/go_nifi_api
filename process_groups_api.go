package go_nifi_api

import (
	"encoding/json"
	"fmt"

	"./types"
)

func (a *app) ProcessGroups(GroupId, token, method string) *types.ProcessGroupEntity {

	processGroupEntity := types.ProcessGroupEntity{}

	url := fmt.Sprintf("%s/%s/%s", a.host, "process-groups", GroupId)
	bytes, _ := a.Do(url, token, method, nil)
	_ = json.Unmarshal(bytes, &processGroupEntity)

	return &processGroupEntity

}
