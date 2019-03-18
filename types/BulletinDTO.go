package main

type BulletinDTO struct {
	Id          int64  `json:"id"`
	NodeAddress string `json:"nodeaddress"`
	Category    string `json:"category"`
	GroupId     string `json:"groupid"`
	SourceId    string `json:"sourceid"`
	SourceName  string `json:"sourcename"`
	Level       string `json:"level"`
	Message     string `json:"message"`
	Timestamp   string `json:"TIMESTAMP"`
}
