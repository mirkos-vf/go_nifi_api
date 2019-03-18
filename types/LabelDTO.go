package types

type LabelDTO struct {
	Id                   string      `json:"id"`
	VersionedComponentId string      `json:"versionedcomponentid"`
	ParentGroupId        string      `json:"parentgroupid"`
	Position             PositionDTO `json:"position"`
	Label                string      `json:"label"`
	Width                float32     `json:"width"`
	Height               float32     `json:"height"`
	Style                object      `json:"style"`
}
