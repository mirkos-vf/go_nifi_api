package go_nifi_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go_nifi_api/types"
)

//"go_nifi_api/types"

func (a *app) ProcessGroups(GroupId, token string) *types.ProcessGroupEntity {

	processGroupEntity := types.ProcessGroupEntity{}

	url := fmt.Sprintf("%s/%s/%s", a.host, "nifi-api/process-groups", GroupId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "*/*")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := a.client.Do(req)
	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	bytes := []byte(data)
	_ = json.Unmarshal(bytes, &processGroupEntity)

	defer response.Body.Close()

	return &processGroupEntity

}
