package types

type ProcessorDTO struct {
	Id                         string            `json:"id"`
	VersionedComponentId       string            `json:"versionedcomponentid"`
	ParentGroupId              string            `json:"parentgroupid"`
	Position                   PositionDTO       `json:"position"`
	Name                       string            `json:"name"`
	Type                       string            `json:"type"`
	Bundle                     BundleDTO         `json:"bundle"`
	State                      string            `json:"state"`
	Style                      object            `json:"style"`
	Relationships              []RelationshipDTO `json:"relationships"`
	Description                string            `json:"description"`
	SupportsParallelProcessing bool              `json:"supportsparallelprocessing"`
	SupportsEventDriven        bool              `json:"supportseventdriven"`
	SupportsBatching           bool              `json:"supportsbatching"`
	PersistsState              bool              `json:"persistsstate"`
	InputRequirement           string            `json:"inputrequirement"`
	Config                     string            `json:"config"`
	ValidationErrors           []string          `json:"validationerrors"`
}
