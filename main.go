package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"./types"
)

const (
	api_url = "https://vm-pd-nifi-1.dh.rt.ru:8080/nifi-api"
)

type app struct {
	Token string
}

type client struct {
	client *http.Client
}

func (c *client) makeClient() {
	// Create New http Transport
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	// Create Http Client
	c.client = &http.Client{Transport: transCfg}
}

func main() {

	a := app{}

	c := client{}
	c.makeClient()

	form := url.Values{}
	form.Set("username", "vladislav.mirkos")
	form.Set("password", "Gfhjkmxxxigrik1234")

	req, err := http.NewRequest("POST", "https://vm-pd-nifi-1.dh.rt.ru:8080/nifi-api/access/token", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	// Check Error
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	a.Token = string(data)
	fmt.Println(a.Token)
	//------------------
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}
	//--------------------------------------------------

	processGroupEntity := types.ProcessGroupEntity{}

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
	bytes := []byte(data1)
	_ = json.Unmarshal(bytes, &processGroupEntity)
	// Check Error
	if err3 != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(data1))

	defer resp1.Body.Close()

}
