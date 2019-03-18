package types

type UserDTO struct {
	Id                   string                      `json:"id"`
	VersionedComponentId string                      `json:"versionedcomponentid"`
	ParentGroupId        string                      `json:"permissionsdto"`
	Position             PositionDTO                 `json:"position"`
	Identity             string                      `json:"identity"`
	Configurable         bool                        `json:"configurable"`
	UserGroups           []TenantEntity              `json:"usergroups"`
	AccessPolicies       []AccessPolicySummaryEntity `json:"accesspolicies"`
}
