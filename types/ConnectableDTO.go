package main

type ConnectableDTO struct {
	Id                   string `json:"id"`
	VersionedComponentId string `json:"versionedcomponentid"`
	Type                 string `json:"type"`
	GroupId              string `json:"groupid"`
	Name                 string `json:"name"`
	Running              bool   `json:"running"`
	Transmitting         bool   `json:"transmitting"`
	Exists               bool   `json:"exists"`
	Comments             string `json:"сomments"`
}
