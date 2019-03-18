package main

type PortDTO struct {
	Id                               string      `json:"id"`
	VersionedComponentId             string      `json:"versionedcomponentid"`
	ParentGroupId                    string      `json:"parentgroupid"`
	Position                         PositionDTO `json:"position"`
	Name                             string      `json:"name"`
	Comments                         string      `json:"comments"`
	State                            string      `json:"state"`
	Type                             string      `json:"type"`
	Transmitting                     bool        `json:"transmitting"`
	ConcurrentlySchedulableTaskCount int32       `json:"concurrentlyschedulabletaskcount"`
	UserAccessControl                []string    `json:"useraccesscontrol"`
	GroupAccessControl               []string    `json:"groupaccesscontrol"`
	ValidationErrors                 []string    `json:"validationerrors"`
}
