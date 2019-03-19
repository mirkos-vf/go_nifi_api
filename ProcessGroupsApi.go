package go_nifi_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go_nifi_api/types"
)

//"go_nifi_api/types"

func (a *app) ProcessGroups(GroupId, token string) *types.ControllerServiceEntity {

	controllerServiceEntity := types.ControllerServiceEntity{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", a.host, "process-groups", GroupId), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, _ := a.client.Do(req)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	bytes := []byte(data)
	_ = json.Unmarshal(bytes, &controllerServiceEntity)

	return &controllerServiceEntity

}

/*

req1, err1 := http.NewRequest("GET", "https://vm-pd-nifi-1.dh.rt.ru:8080/nifi-api/process-groups/518182e7-0168-1000-ffff-ffff8966214b", nil)
if err1 != nil {
log.Fatalln(err)
}
req1.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))

resp1, err2 := c.client.Do(req1)
if err2 != nil {
log.Fatalln(err)
}

data1, err3 := ioutil.ReadAll(resp1.Body)
// Check Error
if err3 != nil {
fmt.Println(err)
os.Exit(1)
}

fmt.Println(string(data1))

defer resp1.Body.Close()

*/
