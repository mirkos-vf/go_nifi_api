package main

type VersionControlInformationDTO struct {
	GroupId          string `json:"groupid"`
	RegistryId       string `json:"registryid"`
	RegistryName     string `json:"registryname"`
	BucketId         string `json:"bucketid"`
	BucketName       string `json:"bucketname"`
	FlowId           string `json:"flowid"`
	FlowName         string `json:"flowname"`
	FlowDescription  string `json:"flowdescription"`
	Version          int32  `json:"version"`
	State            string `json:"state"`
	StateExplanation string `json:"stateexplanation"`
}
