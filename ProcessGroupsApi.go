package go_nifi_api

import (
	"encoding/json"
	"fmt"

	"go_nifi_api/types"
)

//"go_nifi_api/types"

func (a *app) ProcessGroups(GroupId, token, mrthod string) *types.ProcessGroupEntity {

	processGroupEntity := types.ProcessGroupEntity{}

	url := fmt.Sprintf("%s/%s/%s", a.host, "process-groups", GroupId)
	bytes := a.Do(url, token, mrthod, nil)
	_ = json.Unmarshal(bytes, &processGroupEntity)

	return &processGroupEntity

}
