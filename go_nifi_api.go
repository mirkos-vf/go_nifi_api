// Getting Started
//
//package main
//
//import (
//	"fmt"
//	"github.com/mirkos-vf/go_nifi_api"
//	"os"
//)
//
//func main() {
//	nifi := go_nifi_api.NewNiFi("https://localhost:8080")
//
//	token, _ := nifi.AccessToken(os.Getenv("NIFI_USERNAME"), os.Getenv("NIFI_PASSWORD"))
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(token)
//
//}
//
package go_nifi_api

//go:generate python ./generateApi.py -i apis.json
