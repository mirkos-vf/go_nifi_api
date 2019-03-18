package main

type FunnelDTO struct {
	Id                   string      `json:"id"`
	VersionedComponentId string      `json:"versionedcomponentid"`
	ParentGroupId        string      `json:"parentgroupid"`
	Position             PositionDTO `json:"position"`
}
