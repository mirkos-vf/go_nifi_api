package main

type RevisionDTO struct {
	ClientId     string `json:"clientid"`
	Version      int    `json:"version"`
	LastModifier string `json:"lastmodifier"`
}
